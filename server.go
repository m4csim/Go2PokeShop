package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/m4csim/Go2PokeShop/data"
	"github.com/m4csim/Go2PokeShop/mongoconnect"
	"github.com/m4csim/Go2PokeShop/req"
)

const port = ":5500"

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/pokemons", pokemons).Methods("GET")
	router.HandleFunc("/restock", restock).Methods("GET")
	router.HandleFunc("/dbcheck", check_status_db).Methods("GET")
	router.HandleFunc("/dropcoll", drop_collec).Methods("GET")

	fmt.Println("Serving @ http://127.0.0.1" + port)
	log.Fatal(http.ListenAndServe(port, router))
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is not the root page"))
}

func lego() (result data.StockPokemon, err error) {
	err = req.Do("pokemon/132", &result.Pokemon)
	return result, err
}

func check_status_db(w http.ResponseWriter, r *http.Request) {
	mongoconnect.CheckConnect()
	w.Write([]byte("DB seem's ok"))
}

func drop_collec(w http.ResponseWriter, r *http.Request) {
	mongoconnect.Recreate_db()
	w.Write([]byte("Pokeshop collec dropped"))
}

func pokemons(w http.ResponseWriter, r *http.Request) {
	res, err := lego()

	b, err := json.Marshal(res)
	print(b)
	if err != nil {
		fmt.Println("Error:", err)
	}
	w.Write([]byte(b))
}

func restock(w http.ResponseWriter, r *http.Request) {
	mongoconnect.Fixtures_db()
	w.Write([]byte("Restock is done"))
}
