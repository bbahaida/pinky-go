package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"pinky-go/cmd/pinky/product"
)

func main() {
	// Set up the router
	r := mux.NewRouter()

	product.Register(r)
	// Start the server
	http.ListenAndServe(":8080", r)
}
