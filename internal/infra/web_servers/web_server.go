package web_servers

import (
	"github.com/luizFaleiros/GoCommerce/internal/infra/database"
	"github.com/luizFaleiros/GoCommerce/internal/infra/web_servers/handlers"
	"gorm.io/gorm"
	"net/http"
)

func NewWebServer(db *gorm.DB) {
	productDb := database.NewProductDb(db)
	productHandler := handlers.NewProductHander(productDb)
	http.HandleFunc("/products", productHandler.CreateProduct)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
		return
	}
}
