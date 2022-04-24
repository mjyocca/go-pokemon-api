package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getPostgresUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}

// our initial migration function
func initialMigration() *gorm.DB {
	dbURL := getPostgresUrl()
	fmt.Println(dbURL)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	db.AutoMigrate(&Pokemon{})
	return db
}

func handleRequests() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})
	r.HandleFunc("/pokemon", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "All 150 pokemons!")
	})
	r.HandleFunc("/pokemon/{name}", GetPokemon).Methods("GET")
	r.HandleFunc("/pokemon/{name}", UpdatePokemon).Methods("PUT")
	return r
}

func main() {
	initialMigration()
	router := handleRequests()
	fmt.Println("routes registered")
	http.ListenAndServe(":8080", router)
}
