package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go-book-store/pkg/routes"
)

// create a server
// tell where the routes are 

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)

	log.Fatal(http.ListenAndServe(":8080",r))
}