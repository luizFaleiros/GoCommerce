package entity_test

import (
	"github.com/luizFaleiros/GoCommerce/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_NewUser(t *testing.T) {
	user, err := entity.NewUser("nome", "email", "senha")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "nome", user.Name)
	assert.Equal(t, "email", user.Email)
	assert.NotEmpty(t, user.Password)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := entity.NewUser("nome", "email", "senha")
	assert.Nil(t, err)
	assert.True(t, user.ValidatPassword("senha"))
	assert.NotEqual(t, "senha", user.Password)
}
