package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/luizFaleiros/GoCommerce/internal/dto"
	"github.com/luizFaleiros/GoCommerce/internal/entity"
	"github.com/luizFaleiros/GoCommerce/internal/infra/database"
	entityPkg "github.com/luizFaleiros/GoCommerce/pkg/entity"
	"net/http"
	"strconv"
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

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, err := entityPkg.ParseID(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProductDb.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, err := entityPkg.ParseID(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newProduct *entity.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	update, err := h.ProductDb.Update(id, newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(update)
}

func (h ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, err := entityPkg.ParseID(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.ProductDb.DeleteById(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h ProductHandler) FindAllProducts(w http.ResponseWriter, r *http.Request) {
	pages, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		pages = 0
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}
	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDb.FindAll(pages, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
