package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/luizFaleiros/GoCommerce/internal/dto"
	"github.com/luizFaleiros/GoCommerce/internal/entity"
	"github.com/luizFaleiros/GoCommerce/internal/exceptions"
	"github.com/luizFaleiros/GoCommerce/internal/infra/database"
	entityPkg "github.com/luizFaleiros/GoCommerce/pkg/entity"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductDb database.ProductInterface
}

func NewProductHander(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDb: db}
}

// Login godoc
// @Sumary Create Product
// @Description Endpoint que faz o login do usuario
// @Tags products
// @Accept json
// @Produce json
// @Param  request body dto.CreateProductDTO true "product create"
// @Success 201
// @Failure 400 {object} exceptions.Error
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.CreateProductDTO
	if err := json.NewDecoder(r.Body).Decode(&productDto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	var p *entity.Product
	if product, err := entity.NewProduct(productDto.Name, productDto.Price); err == nil {
		p = product
	} else {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if err := h.ProductDb.Create(p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Login godoc
// @Sumary Create Product
// @Description Endpoint que faz o login do usuario
// @Tags products
// @Accept json
// @Produce json
// @Param  id path string true "product ID" Format(uuid)
// @Success 200 {object} entity.Product
// @Failure 400 {object} exceptions.Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{"Id not found"}
		json.NewEncoder(w).Encode(error)
		return
	}
	if _, err := entityPkg.ParseID(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	product, err := h.ProductDb.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := exceptions.Error{err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {

}

// Login godoc
// @Sumary Create Product
// @Description Endpoint que faz o login do usuario
// @Tags products
// @Accept json
// @Produce json
// @Param  id path string true "product ID" Format(uuid)
// @Success 200 {object} entity.Product
// @Failure 400 {object} exceptions.Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{"Token not found"}
		json.NewEncoder(w).Encode(error)
		return
	}
	if _, err := entityPkg.ParseID(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	var newProduct *entity.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	update, err := h.ProductDb.Update(id, newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(update)
}

// Login godoc
// @Sumary Create Product
// @Description Endpoint que faz o login do usuario
// @Tags products
// @Accept json
// @Produce json
// @Param  id path string true "product ID" Format(uuid)
// @Success 200
// @Failure 404 {object} exceptions.Error
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{"Token not found"}
		json.NewEncoder(w).Encode(error)
		return
	}
	if _, err := entityPkg.ParseID(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	if err := h.ProductDb.DeleteById(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := exceptions.Error{err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
}

// Login godoc
// @Sumary Create Product
// @Description Endpoint que faz o login do usuario
// @Tags products
// @Accept json
// @Produce json
// @Param  page query string false "Page number"
// @Param  limit query string false "limit number"
// @Param  sort query string false "asc,desc"
// @Success 200 {array} entity.Product
// @Failure 400 {object} exceptions.Error
// @Router /products [get]
// @Security ApiKeyAuth
func (h ProductHandler) FindAllProducts(w http.ResponseWriter, r *http.Request) {
	pages, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		pages = 0
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}
	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDb.FindAll(pages, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := exceptions.Error{err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
