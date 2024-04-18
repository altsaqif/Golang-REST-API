package routes

import (
	"github.com/altsaqif/go-restapi-mux/cmd/controllers/authController"
	"github.com/altsaqif/go-restapi-mux/cmd/controllers/homeController"
	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()
	router.HandleFunc("/register", authController.Register).Methods("POST")
	router.HandleFunc("/login", authController.Login).Methods("POST")
	router.HandleFunc("/logout", authController.Logout).Methods("GET")
	router.HandleFunc("/", homeController.Home).Methods("GET")
}
