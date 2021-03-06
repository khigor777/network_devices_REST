package main

import (
	"log"
	"net/http"
)

const port  = ":8080"

var devices Devices

func handleRequests() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(port, router))
}

func main() {
	devices = localDevices()
	handleRequests()
}
