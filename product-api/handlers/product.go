package handlers

import (
	"log"
	"net/http"
	"product/product-api/data"
)

type Products struct {
	log *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{log: logger}
}

func (p *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		p.getProducts(writer, request)
		return
	}
	if request.Method == http.MethodPut {

	}

	writer.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(w)
	if err != nil {
		http.Error(w, "Unable to decode json", http.StatusInternalServerError)
	}
}
