package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (p *Products) DeleteProduct(writter http.ResponseWriter, r *http.Request) {
	p.log("Trying to delete a product")
	vars := mux.Vars(r)
	id := strconv.Atoi(vars["id"])
	p.
}
