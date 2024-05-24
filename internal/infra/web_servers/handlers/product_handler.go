package handlers

import (
	"encoding/json"
	"github.com/luizFaleiros/GoCommerce/internal/dto"
	"github.com/luizFaleiros/GoCommerce/internal/entity"
	"github.com/luizFaleiros/GoCommerce/internal/infra/database"
	"net/http"
)

type ProductHandler struct {
	ProductDb database.ProductInterface
}

func NewProductHander(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDb: db}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.CreateProductDTO

	if err := json.NewDecoder(r.Body).Decode(&productDto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var p *entity.Product
	if product, err := entity.NewProduct(productDto.Name, productDto.Price); err == nil {
		p = product
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.ProductDb.Create(p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
