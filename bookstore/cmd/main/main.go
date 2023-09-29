package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamvibs/go-projects/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

// change controller to pkglication
