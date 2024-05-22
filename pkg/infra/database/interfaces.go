package database

import "github.com/luizFaleiros/GoCommerce/internal/entity"

type UserInterface interface {
	Create(user *entity.User)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
}
