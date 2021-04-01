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
