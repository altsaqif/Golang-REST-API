package routes

import (
	"github.com/altsaqif/go-restapi-mux/cmd/controllers/productController"
	"github.com/altsaqif/go-restapi-mux/cmd/middlewares"
	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	router := r.PathPrefix("/products").Subrouter()

	// Penerapan Midlleware
	router.Use(middlewares.Auth)

	router.HandleFunc("/product", productController.Products).Methods("GET")
	router.HandleFunc("/product/{id}", productController.Product).Methods("GET")
	router.HandleFunc("/product", productController.Create).Methods("POST")
	router.HandleFunc("/product/{id}", productController.Update).Methods("PUT")
	router.HandleFunc("/product", productController.Delete).Methods("DELETE")
}
