package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model
	Name  string `json:"name" gorm:"primaryKey"`
	Type  string `json:"type"`
	Moves string `json:"moves"`
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Get Pokemon: %s", name)
}

func CreatePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreatePokemon")
	fmt.Fprintf(w, "Creating pokemon")
}

func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdatePokemon")
	fmt.Fprintf(w, "Updating pokemon")
}
