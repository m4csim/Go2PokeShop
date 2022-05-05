package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":5500"

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/pokemons", pokemons).Methods("GET")

	fmt.Println("Serving @ http://127.0.0.1" + port)
	log.Fatal(http.ListenAndServe(port, router))
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is not the root page"))
}

func pokemons(w http.ResponseWriter, r *http.Request) {

	// resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	var pokemonList = []data.pokemon{
		data.pokemon{1, "Terre", "Bulbasaur", 45, 49, "Mist-Ball", 50.00, 4},
		data.pokemon{2, "Terre", "Ivysaur", 60, 62, "Psychoboost", 50.00, 4},
		data.pokemon{3, "Terre", "Venusaur", 80, 82, "Overheat", 50.00, 4},
	}

	b, err := json.Marshal(pokemonList)
	print(b)
	if err != nil {
		fmt.Println("Error:", err)
	}
	w.Write([]byte(b))
}
