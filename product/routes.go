package product

import "net/http"

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/products" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		GetProductsHandler(w, r)
	case http.MethodPost:
		CreateProductHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ProductsByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetProductByIDHandler(w, r)
	case http.MethodDelete:
		DeleteProductByIDHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
