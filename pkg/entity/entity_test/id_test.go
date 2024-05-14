package entity_test

import (
	"github.com/luizFaleiros/GoCommerce/pkg/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewID(t *testing.T) {
	id := entity.NewID()
	assert.NotEmpty(t, id)
}

func TestParseID(t *testing.T) {
	id, err := entity.ParseID(entity.NewID().String())
	assert.Nil(t, err)
	assert.NotEmpty(t, id)
}
