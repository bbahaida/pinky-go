package product

import (
	"github.com/gorilla/mux"
	"net/http"
	"pinky-go/internal/api"
	"pinky-go/internal/service"
)

func Register(r *mux.Router) {
	// Initialize ProductService
	productService := service.NewProductService() // Assuming a concrete implementation exists

	// Create a new instance of ProductHandler
	productHandler := api.NewProductHandler(productService)
	r.HandleFunc("/products", productHandler.GetProducts).Methods(http.MethodGet)
	r.HandleFunc("/products", productHandler.CreateProduct).Methods(http.MethodPost)
}
