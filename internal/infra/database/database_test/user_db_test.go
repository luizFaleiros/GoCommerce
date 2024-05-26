package database_test

import (
	"github.com/luizFaleiros/GoCommerce/internal/entity"
	"github.com/luizFaleiros/GoCommerce/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	user, _ := entity.NewUser("test", "test@test", "test")
	db.AutoMigrate(&entity.User{})

	userDb := database.NewUserDb(db)
	err = userDb.Create(user)
	var createdUser entity.User
	error := db.First(&createdUser, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Nil(t, error)
	assert.NotNil(t, createdUser)
	assert.Equal(t, user.ID, createdUser.ID)
	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)

}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	user, _ := entity.NewUser("test", "test@test", "test")
	db.AutoMigrate(&entity.User{})
	userDb := database.NewUserDb(db)
	err = userDb.Create(user)
	if err != nil {
		t.Error(err)
	}
	createdUser, error := userDb.FindByEmail(user.Email)
	assert.Nil(t, error)
	assert.NotNil(t, createdUser)
	assert.Equal(t, user.ID, createdUser.ID)
	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)
}
