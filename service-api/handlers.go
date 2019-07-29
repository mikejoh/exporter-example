package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Not implemented!")
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(info)
}

func incItems(w http.ResponseWriter, r *http.Request) {
	var itemIncrement Increment

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please add an increment")
	}

	err = json.Unmarshal(reqBody, &itemIncrement)
	if err != nil {
		fmt.Fprintf(w, "Failed when decoding JSON")
	}

	info.NumItems += itemIncrement.Number

	fmt.Fprintf(w, "Incremented items with %d", itemIncrement.Number)
}