package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	info = &Info{}
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/api/info", getInfo).Methods("GET")
	router.HandleFunc("/api/items", incItems).Methods("POST")

	log.Println("Starting Service API!")
	log.Fatal(http.ListenAndServe(":8000", router))
}
