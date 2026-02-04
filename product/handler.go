package product

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := GetProducts()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var input CreateProductInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	product := CreateProduct(input)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	product, found := GetProductByID(id)
	if !found {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
