package entity_test

import (
	"github.com/google/uuid"
	"github.com/luizFaleiros/GoCommerce/internal/entity"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestProduct_Validate(t *testing.T) {
	tests := []struct {
		name      string
		product   entity.Product
		wantError error
	}{
		{
			name: "missing name",
			product: entity.Product{
				ID:        uuid.New(),
				Name:      "",
				Price:     100.0,
				CreatedAt: time.Now(),
			},
			wantError: entity.ErrNameIsRequired,
		},
		{
			name: "price is zero",
			product: entity.Product{
				ID:        uuid.New(),
				Name:      "Valid Name",
				Price:     0.0,
				CreatedAt: time.Now(),
			},
			wantError: entity.ErrPriceIsRequired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.product.Validate()
			assert.Equal(t, tt.wantError, err)
		})
	}
}

func TestNewProductValid(t *testing.T) {
	productTest := struct {
		name         string
		productName  string
		productPrice float64
		wantError    error
	}{
		name:         "Valid product",
		productName:  randomString(100),
		productPrice: randomPrice(),
		wantError:    nil,
	}
	t.Run(productTest.name, func(t *testing.T) {
		product, err := entity.NewProduct(productTest.productName, productTest.productPrice)
		assert.Equal(t, productTest.wantError, err)
		assert.Equal(t, productTest.productName, product.Name)
		assert.Equal(t, productTest.productPrice, product.Price)
		assert.NotNil(t, product.ID)
		assert.NotNil(t, product.CreatedAt)
	})
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomPrice() float64 {
	return 0.01 + rand.Float64()*1000
}
