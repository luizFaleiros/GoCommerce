package database

import "github.com/luizFaleiros/GoCommerce/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindById(id string) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]*entity.Product, error)
	Update(id string, newProduct *entity.Product) (*entity.Product, error)
	DeleteById(id string) error
}
