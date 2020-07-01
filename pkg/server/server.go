package main

import (
	"log"
	"net/http"
	"self-study-golang/pkg/server/api"
)

func main() {
	handleRequest()
}

func handleRequest() {
	http.Handle("/", api.HomepageHandler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

