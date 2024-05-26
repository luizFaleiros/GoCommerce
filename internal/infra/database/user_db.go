package database

import (
	"github.com/luizFaleiros/GoCommerce/internal/entity"
	"gorm.io/gorm"
)

type UserDb struct {
	DB *gorm.DB
}

func NewUserDb(db *gorm.DB) *UserDb {
	return &UserDb{DB: db}
}

func (u *UserDb) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserDb) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDb) Update(user *entity.User) (*entity.User, error) {
	oldUser, err := u.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	user.ID = oldUser.ID
	if err := u.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
