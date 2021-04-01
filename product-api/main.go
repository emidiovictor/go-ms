package main

import (
	"context"
	"go-ms/product-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	produHandler := handlers.NewProducts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", produHandler.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", produHandler.UpdateProduct)
	putRouter.Use(produHandler.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", produHandler.AddProduct)
	postRouter.Use(produHandler.MiddlewareProductValidation)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)

	// sm.Handle("/products", produHandler)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Println("BA DUM TIS")
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println("Down gracefully", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
