package product

import "net/http"

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetProductsHandler(w, r)
	case http.MethodPost:
		CreateProductHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
