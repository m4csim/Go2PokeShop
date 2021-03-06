package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	// "github.com/m4csim/Go2PokeShop/data"
	"github.com/m4csim/Go2PokeShop/data"
	"github.com/m4csim/Go2PokeShop/database"
	"github.com/m4csim/Go2PokeShop/req"
	// "github.com/m4csim/Go2PokeShop/req"
)

const port = ":5500"

func main() {
	// commandes
	// pokemons à gérer
	//
	//connect
	database.ConnectDB()
	// credentials := handlers.AllowCredentials()
	// methods := handlers.AllowedMethods([]string{"POST"})
	// origins := handlers.AllowedOrigins([]string{"www.example.com"})

	router := mux.NewRouter()
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/pokemons", all_pokemons).Methods("GET")
	router.HandleFunc("/pokemons/{pokemon}", pokemons).Methods("GET")
	router.HandleFunc("/pokemons/{pokemon}/{number}", buy_pokemons).Methods("GET")
	router.HandleFunc("/restock", restock).Methods("GET")
	router.HandleFunc("/dropcoll", drop_collec).Methods("GET")
	handler := cors.Default().Handler(router)
	fmt.Println("Serving @ http://127.0.0.1" + port)
	log.Fatal(http.ListenAndServe(port, handler))
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is not the root page"))
}

func drop_collec(w http.ResponseWriter, r *http.Request) {
	database.Recreate_db()
	w.Write([]byte("Pokeshop collec dropped"))
}

func pokemons(w http.ResponseWriter, r *http.Request) {
	pokemon := mux.Vars(r)["pokemon"]
	minified := r.URL.Query().Get("minified")
	result := database.Get_one_pokemon(pokemon)
	if result.Name != "" {
		if minified == "1" {

			var pokemon data.MinifiedPokemon

			req.Do(fmt.Sprintf("pokemon/%s", result.Name), &pokemon)
			json_response, json_err := json.Marshal(pokemon)
			if json_err != nil {
				w.Write([]byte("error"))
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(json_response)
		} else {
			var pokemon data.DetailedPokemon
			req.Do(fmt.Sprintf("pokemon/%s", result.Name), &pokemon)
			json_response, json_err := json.Marshal(pokemon)
			if json_err != nil {
				w.Write([]byte("error"))
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(json_response)
		}
	} else {
		w.Write([]byte("not found"))
	}
}

func buy_pokemons(w http.ResponseWriter, r *http.Request) {
	pokemon := mux.Vars(r)["pokemon"]
	number := mux.Vars(r)["number"]
	result := database.Get_one_pokemon(pokemon)
	fmt.Println(result)
	if result.Name != "" {
		if res, _ := strconv.Atoi(number); result.Count > res {
			new_count := result.Count - res
			res := database.Destock_pokemon(pokemon, new_count)
			fmt.Println(res)
			if res != nil {
				w.Write([]byte("true"))
			} else {
				w.Write([]byte("false"))
			}
		} else {
			w.Write([]byte("cannot buy that much"))
		}
	} else {
		w.Write([]byte("not found"))
	}
}

func all_pokemons(w http.ResponseWriter, r *http.Request) {
	result := database.Get_all_pokemon()
	json_response, json_err := json.Marshal(result)
	if json_err != nil {
		w.Write([]byte("error"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_response)
}

func restock(w http.ResponseWriter, r *http.Request) {
	database.Fixtures_db()
	w.Write([]byte("Restock is done"))
}
