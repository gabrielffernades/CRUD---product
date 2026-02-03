package product

import (
	"encoding/json"
	"net/http"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products := GetProducts()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
