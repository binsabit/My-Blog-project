package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(home))
	mux.Get("/profile", http.HandlerFunc(profile))
	mux.Get("/store", http.HandlerFunc(showStore))
	mux.Get("/store/:id", http.HandlerFunc(showStoreItem))

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
