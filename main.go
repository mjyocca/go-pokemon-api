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
	r.HandleFunc("/pokemon", GetAllPokemon).Methods("GET")
	r.HandleFunc("/pokemon/{name}", GetPokemon).Methods("GET")
	r.HandleFunc("/pokemon/{name}", UpdatePokemon).Methods("PUT")
	r.HandleFunc("/pokemon", CreatePokemon).Methods("POST")
	return r
}

func main() {
	// setup db and handle migrations
	InitDatabase()
	http.ListenAndServe(":8080", handleRequests())
	fmt.Println("routes registered")
}
