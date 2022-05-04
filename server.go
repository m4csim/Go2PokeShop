package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":5500"

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/pokemons/", pokemons).Methods("GET")

	fmt.Println("Serving @ http://127.0.0.1" + port)
	log.Fatal(http.ListenAndServe(port, router))
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is not the root page"))
}

func pokemons(w http.ResponseWriter, r *http.Request) {

	type pokemon struct {
		Id      int
		Type    string
		Name    string
		Hpoint  int
		Pattack int
		Smove   string
		Price   float64
		Count   int
	}

	var pokemonList = []pokemon{
		pokemon{1, "Terre", "Bulbasaur", 45, 49, "Mist-Ball", 50.00, 4},
		pokemon{2, "Terre", "Ivysaur", 60, 62, "Psychoboost", 50.00, 4},
		pokemon{3, "Terre", "Venusaur", 80, 82, "Overheat", 50.00, 4},
	}
}
