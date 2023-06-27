package api

import (
	"encoding/json"
	"net/http"

	"github.com/muhangga/repository"
)

type ProductListErrorResponse struct {
	Error string `json:"error"`
}

type ProductListSuccessResponse struct {
	Products []repository.Product `json:"products"`
}

func (api *API) productList(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	resp := ProductListSuccessResponse{}
	resp.Products = make([]repository.Product, 0)

	products, err := api.productRepo.FetchProduct()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(ProductListErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	for _, product := range products {
		resp.Products = append(resp.Products, repository.Product{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	w.WriteHeader(http.StatusOK)
	encoder.Encode(resp)
}
