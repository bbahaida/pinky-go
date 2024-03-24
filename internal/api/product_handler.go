package api

import (
	"encoding/json"
	"net/http"
	"pinky-go/internal/model"
	"pinky-go/internal/service"
)

const CONTENT_TYPE = "application/json"

type ProductHandler struct {
	productService service.ProductService
}

// NewProductHandler creates a new ProductHandler with the given ProductService.
func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// GetProducts handles GET requests to fetch all products.
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.GetAllProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", CONTENT_TYPE)
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct model.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdProduct, err := h.productService.CreateProduct(r.Context(), newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", CONTENT_TYPE)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}
