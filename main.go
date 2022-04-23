package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})
	r.HandleFunc("/pokemon", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "All 150 pokemon")
	})
	r.HandleFunc("/pokemon/{name}", GetPokemon).Methods("GET")
	r.HandleFunc("/pokemon/{name}", UpdatePokemon).Methods("PUT")
	return r
}

func main() {
	router := handleRequests()
	http.ListenAndServe(":80", router)
}
