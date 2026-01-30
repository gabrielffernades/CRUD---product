package main

import (
	"net/http"

	"github.com/gabrielffernandes/CRUD---product/handlers"
	"github.com/gabrielffernandes/CRUD---product/product"
)

func main() {
	http.HandleFunc("/health", handlers.HealthHandler)

	http.HandleFunc("/products", product.GetProductsHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
