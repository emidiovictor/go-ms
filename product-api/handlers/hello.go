package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{log: logger}
}

func (h *Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.log.Print("Hello World")
	d, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "oops", http.StatusBadRequest)
		return
	}
	h.log.Printf("Data %s \n", d)
}
