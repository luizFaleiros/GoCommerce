package handlers

import (
	"encoding/json"
	"github.com/go-chi/jwtauth/v5"
	"github.com/luizFaleiros/GoCommerce/internal/dto"
	"github.com/luizFaleiros/GoCommerce/internal/entity"
	"github.com/luizFaleiros/GoCommerce/internal/exceptions"
	"github.com/luizFaleiros/GoCommerce/internal/infra/database"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	UserDb    database.UserInterface
	Jwt       *jwtauth.JWTAuth
	JwtExpiry int64
}

func NewUserHandler(db database.UserInterface, Jwt *jwtauth.JWTAuth, JwtExpiry int64) *UserHandler {
	return &UserHandler{db, Jwt, JwtExpiry}
}

// Create User godoc
// @Sumary Create user
// @Description End point para criar usuarios
// @Tags users
// @Accept json
// @Produce json
// @Param  request body dto.CreateUserDTO true "user request"
// @Success 201
// @Failure 400 {object} exceptions.Error
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.UserDb.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Login godoc
// @Sumary Login
// @Description Endpoint que faz o login do usuario
// @Tags users
// @Accept json
// @Produce json
// @Param  request body dto.LoginDTO true "user credentials"
// @Success 200 {object} dto.TokenDTO
// @Failure 400 {object} exceptions.Error
// @Failure 401 {object} exceptions.Error
// @Failure 404 {object} exceptions.Error
// @Router /users/login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userDto dto.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	user, err := h.UserDb.FindByEmail(userDto.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := exceptions.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if !user.ValidatPassword(userDto.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		error := exceptions.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	_, token, err := h.Jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Hour * time.Duration(h.JwtExpiry)).Unix(),
	})
	if err != nil {
		log.Fatalln(err)
	}
	accessToken := dto.TokenDTO{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
	w.WriteHeader(http.StatusOK)
}
