package database_test

import (
	"fmt"
	"github.com/luizFaleiros/GoCommerce/internal/entity"
	"github.com/luizFaleiros/GoCommerce/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	product, _ := entity.NewProduct("test", 1.60)
	db.AutoMigrate(&entity.Product{})

	productDb := database.NewProductDb(db)
	err = productDb.Create(product)
	var createdProduct entity.Product
	error := db.First(&createdProduct, "id = ?", product.ID).Error
	assert.Nil(t, err)
	assert.Nil(t, error)
	assert.NotNil(t, createdProduct)
	assert.Equal(t, product.ID, createdProduct.ID)
	assert.Equal(t, product.Name, createdProduct.Name)
	assert.Equal(t, product.Price, createdProduct.Price)
}

func TestFindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	product, _ := entity.NewProduct("test", 1.60)
	db.AutoMigrate(&entity.Product{})

	productDb := database.NewProductDb(db)
	err = productDb.Create(product)

	createdProduct, error := productDb.FindById(product.ID.String())
	assert.Nil(t, error)
	assert.NotNil(t, createdProduct)
	assert.Equal(t, product.ID, createdProduct.ID)
	assert.Equal(t, product.Name, createdProduct.Name)
}

func TestFindAll(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	productDb := database.NewProductDb(db)
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, _ := entity.NewProduct(fmt.Sprintf("Product %d", i), randomPrice())
		assert.NoError(t, err)
		db.Create(&product)
	}
	products, err := productDb.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, len(products), 10)
	assert.Equal(t, products[0].Name, "Product 1")
	assert.Equal(t, products[9].Name, "Product 10")
	products, err = productDb.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, len(products), 10)
	assert.Equal(t, products[0].Name, "Product 11")
	assert.Equal(t, products[9].Name, "Product 20")
	products, err = productDb.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, len(products), 3)
	assert.Equal(t, products[0].Name, "Product 21")
	assert.Equal(t, products[2].Name, "Product 23")

}

func randomPrice() float64 {
	return 0.01 + rand.Float64()*1000
}
