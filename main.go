package main

import (
	"log"
	"net/http"

	. "github.com/altsaqif/go-restapi-gin/config"
	. "github.com/altsaqif/go-restapi-gin/controllers/authController"
	. "github.com/altsaqif/go-restapi-gin/controllers/productController"
	. "github.com/altsaqif/go-restapi-gin/database"
	"github.com/altsaqif/go-restapi-gin/middlewares"
	"github.com/gorilla/mux"
)

func main() {
	ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/register", Register).Methods("POST")
	r.HandleFunc("/logout", Logout).Methods("GET")

	// Penerapan Midlleware
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTMiddleware)

	api.HandleFunc("/products", Index).Methods("GET")
	api.HandleFunc("/products/{id}", Show).Methods("GET")
	api.HandleFunc("/products", Create).Methods("POST")
	api.HandleFunc("/products/{id}", Update).Methods("PUT")
	api.HandleFunc("/products", Delete).Methods("DELETE")

	PORT := GoDotEnvVariable("APP_PORT")
	log.Fatal(http.ListenAndServe(PORT, r))
}
