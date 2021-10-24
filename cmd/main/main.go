package main

import (
	"github.com/gorilla/mux"
	"houseAPI/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterHouseRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost: 9010", r))
}
