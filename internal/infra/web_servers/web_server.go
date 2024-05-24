package web_servers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/luizFaleiros/GoCommerce/internal/infra/database"
	"github.com/luizFaleiros/GoCommerce/internal/infra/web_servers/handlers"
	"gorm.io/gorm"
	"net/http"
)

func NewWebServer(db *gorm.DB) {
	productDb := database.NewProductDb(db)
	productHandler := handlers.NewProductHander(productDb)
	route := chi.NewRouter()
	route.Use(middleware.Logger)
	route.Post("/products", productHandler.CreateProduct)
	route.Get("/products/{id}", productHandler.GetProducts)
	route.Get("/products", productHandler.FindAllProducts)
	route.Put("/products/{id}", productHandler.UpdateProduct)
	route.Delete("/products/{id}", productHandler.DeleteProduct)
	err := http.ListenAndServe(":8000", route)
	if err != nil {
		panic(err)
		return
	}
}
