package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"self-study-golang/pkg/server/api"
)

func main() {
	log.Fatal(run())
}

func run() error {

	r := router()
	http.Handle("/", r)
	return http.ListenAndServe(":8080", nil)
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", api.HomepageHandler{})
	r.HandleFunc("/articles/categories/{category}", api.GetArticleByCategoryHandler)
	return r
}

