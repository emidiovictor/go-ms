package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Print("Hello World")
		d, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "oops", http.StatusBadRequest)
			return
		}
		log.Printf("Data %s \n", d)
	})
	http.ListenAndServe(":8080", nil)
}
