package product

import (
	"encoding/json"
	"net/http"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products := []Product{
		{
			ID:          1,
			Name:        "Teclado gamer",
			Description: "Teclado gamer com iluminação RGB",
			Price:       800.00,
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
