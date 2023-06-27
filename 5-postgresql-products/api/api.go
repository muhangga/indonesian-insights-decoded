package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/muhangga/repository"
)

type API struct {
	productRepo repository.ProductRepository
	mux         *http.ServeMux
}

func NewAPI(productRepo repository.ProductRepository) *API {
	mux := http.NewServeMux()

	api := &API{productRepo, mux}

	mux.Handle("/api/products", api.GET(http.HandlerFunc(api.productList)))

	return api
}

func (api *API) GET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			encoder.Encode(ProductListErrorResponse{Error: "Method not allowed"})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (api *API) Handler() http.Handler {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("Starting API server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
