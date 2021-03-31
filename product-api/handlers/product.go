package handlers

import (
	"go-ms/product-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	if request.Method == http.MethodPost {
		p.addProduct(writer, request)
		return
	}

	if request.Method == http.MethodPut {
		p.log.Printf("PUT")
		r := regexp.MustCompile(`/([0-9]+)`)
		var g = r.FindAllStringSubmatch(request.URL.Path, -1)
		if len(g) != 1 {
			http.Error(writer, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(writer, "Invalid URL", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, _ := strconv.Atoi(idString)
		p.log.Println("GOT ID ", id)
		p.UpdateProduct(id, writer, request)
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
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Trying to post a product")
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.log.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Trying to put a product")
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
