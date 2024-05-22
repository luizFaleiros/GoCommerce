package database

import "github.com/luizFaleiros/GoCommerce/internal/entity"

type UserInterface interface {
	Create(user *entity.User)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindById(id string) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]entity.Product, int, error)
	Update(product *entity.Product) (*entity.Product, error)
	DeleteById(id string) error
}
