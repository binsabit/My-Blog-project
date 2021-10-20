package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Home Page"))
}
func profile(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from Profile Page"))
}
func showStore(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Store Page"))
}
func showStoreItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		fmt.Fprintf(w, "%d", id)

		return
	}

	fmt.Fprintf(w, "Hello from Store Item id = %d", id)
}
