package product

import (
	"encoding/json"
	"net/http"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products := []map[string]interface{}{
		{
			"id":          1,
			"name":        "Teclado gamer",
			"description": "Teclado gamer com iluminação RGB e teclas programáveis",
			"price":       800.00,
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
