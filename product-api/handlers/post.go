package handlers

import (
	"go-ms/product-api/data"
	"net/http"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Trying to post a product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.log.Printf("Prod: %#v", prod)
	data.AddProduct(&prod)
}
