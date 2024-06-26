package entity

import (
	"github.com/luizFaleiros/GoCommerce/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"_"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: string(hash),
		ID:       entity.NewID(),
	}, nil
}

func (u *User) ValidatPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
