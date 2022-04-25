package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model
	ID    string
	Name  string // `gorm:"primaryKey"`
	Type  string
	Moves string
}

type PokemonResponse struct {
	Data   []Pokemon
	Errors []string
}

func GetAllPokemon(w http.ResponseWriter, r *http.Request) {
	var pokemon []Pokemon
	DB.Find(&pokemon)
	fmt.Println("{}", &pokemon)
	json.NewEncoder(w).Encode(pokemon)
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Println("name: {}", name)
	var pokemon Pokemon
	DB.Where(&Pokemon{ID: name}).Find(&pokemon)
	json.NewEncoder(w).Encode(&pokemon)
}

func CreatePokemon(w http.ResponseWriter, r *http.Request) {
	var pokemon Pokemon
	json.NewDecoder(r.Body).Decode(&pokemon)
	DB.Create(&pokemon)
	fmt.Println("{}", &pokemon)
	result := PokemonResponse{
		Data: []Pokemon{
			pokemon,
		},
		Errors: []string{},
	}
	fmt.Println("{}", &pokemon)
	json.NewEncoder(w).Encode(&result)
}

func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var pokemon Pokemon
	json.NewDecoder(r.Body).Decode(&pokemon)
	DB.Where(&Pokemon{ID: name}).Save(&pokemon)
	result := PokemonResponse{
		Data: []Pokemon{
			pokemon,
		},
	}
	json.NewEncoder(w).Encode(&result)
}
