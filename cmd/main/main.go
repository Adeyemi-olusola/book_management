package main

import (
	"book_management/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoresRoutes(r)
	http.Handle("/" , r)
	log.Fatal(http.ListenAndServe("localhost:9010" , r ))
}