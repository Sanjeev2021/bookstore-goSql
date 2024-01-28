package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sanjeev2021/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) // registerbookstorerroutes is in bookstore-routes and we are passing the router
	http.Handle("/", r)
	fmt.Println("The api is running.")
	log.Fatal(http.ListenAndServe("localhost:9010", r)) // helps in creating the server
}
