package main

import (
	"log"
	"net/http"

	"github.com/altsaqif/go-restapi-mux/cmd/configs"
	"github.com/altsaqif/go-restapi-mux/cmd/routes"
	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDatabase()
	r := mux.NewRouter()

	router := r.PathPrefix("/api").Subrouter()
	routes.AuthRoutes(router)
	routes.ProductRoutes(router)

	PORT := configs.GoDotEnvVariable("APP_PORT")
	log.Printf("Connect to http://localhost%s", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}
