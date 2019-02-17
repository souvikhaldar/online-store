package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/souvikhaldar/online-store/pkg/agent"
	"github.com/souvikhaldar/online-store/pkg/locationtracker"
	"github.com/souvikhaldar/online-store/pkg/product"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/product", product.AddProduct).Methods("POST")
	router.HandleFunc("/agent", agent.AddAgent).Methods("POST")
	router.HandleFunc("/purchase", product.Purchase).Methods("GET")
	//router.HandleFunc("/track", locationtracker.Track)
	router.HandleFunc("/longlat", locationtracker.LongLatHandler).Methods("POST")
	router.HandleFunc("/ws", locationtracker.WsHandler)
	go locationtracker.Echo()
	log.Fatal(http.ListenAndServe(":8192", router))
}
