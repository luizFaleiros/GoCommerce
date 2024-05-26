package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/luizFaleiros/GoCommerce/configs"
	_ "github.com/luizFaleiros/GoCommerce/docs"
	"github.com/luizFaleiros/GoCommerce/internal/entity"
	"github.com/luizFaleiros/GoCommerce/internal/infra/database"
	"github.com/luizFaleiros/GoCommerce/internal/infra/web_servers/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

// @title Go ecommerce API
// @Version 1.0
// @description product API
// @termOfService http://swagger.io/terms/

// @contact.name Luiz Faleiros
// @contact.url www.github.com/luizFaleiros
// @contact.email luiz.h.s.faleiros@gmail.com

// @licemse.name open
// @license.url http://www.open.com.br

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	config := configs.LoadConfig()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDb := database.NewProductDb(db)
	productHandler := handlers.NewProductHander(productDb)

	userDb := database.NewUserDb(db)
	userHandler := handlers.NewUserHandler(userDb, config.TokenAuth, config.JWTExperiTime)

	route := chi.NewRouter()
	route.Use(middleware.Logger)
	route.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProducts)
		r.Get("/", productHandler.FindAllProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})
	route.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/login", userHandler.Login)
	})
	route.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	err = http.ListenAndServe(":8000", route)
	if err != nil {
		panic(err)
		return
	}
}
