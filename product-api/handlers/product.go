// Package classification of Product API
//
//	Documentation for Product API
//
//		Schemes: http
//		BasePath: /
//		Version: 1.0.0
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"go-ms/product-api/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//A list of products returns in the response
// swagger:response productResponse
type productResponseWrapper struct {
	//All products in the system
	// in: body
	Body []data.Products
}

type Products struct {
	log *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{log: logger}
}

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
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Trying to post a product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.log.Printf("Prod: %#v", prod)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Impossible to cast the id", http.StatusBadRequest)
		return
	}
	p.log.Println("Trying to PUT a product", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}
		// validate de product

		err = prod.Validate()
		if err != nil {
			p.log.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)

	})
}
