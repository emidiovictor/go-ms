package handlers

import (
	"go-ms/product-api/data"
	"net/http"
)

// swagger:route GET /products products listProduct
// Returns a list of products
// responses :
// 200: productResponse

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(w)
	if err != nil {
		http.Error(w, "Unable to decode json", http.StatusInternalServerError)
	}
}
