package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/souvikhaldar/online-store/pkg/product"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/product", product.AddProduct).Methods("POST")
	log.Fatal(http.ListenAndServe(":8192", router))
}
