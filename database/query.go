package database

import (
	"context"
	//"encoding/json"
	"fmt"
	"github.com/m4csim/Go2PokeShop/data"
	"github.com/m4csim/Go2PokeShop/req"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"time"
)

var DB *mongo.Client

func Recreate_db() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@127.0.0.1:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("pokeshop")
	stocks := database.Collection("stocks")
	stocks.Drop(ctx)
	stocks.InsertOne(ctx, bson.D{})

}

func Fixtures_db() {
	stock_collection := MI.DB.Collection("stocks")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	for i := 0; i < 100; i++ {
		res, _ := generate(rand.Intn(500) + rand.Intn(300) + 1)
		res.Count = rand.Intn(7)
		res.Price = rand.Intn(150)
		//res2B, _ := json.Marshal(res.Pokemon)
		//fmt.Println(string(res2B))
		//b, _ := json.Marshal(res)
		stock_collection.InsertOne(ctx, res)
	}
}

func generate(id int) (result data.StockPokemon, err error) {
	err = req.Do(fmt.Sprintf("pokemon/%d", id), &result)
	return result, err
}

func Get_one_pokemon(name string) data.StockPokemon {
	stock_collection := MI.DB.Collection("stocks")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var pokemon_stock data.StockPokemon
	filter := bson.M{"name": name}
	find_result := stock_collection.FindOne(ctx, filter)
	if err := find_result.Err(); err != nil {
		return data.StockPokemon{}
	}
	err := find_result.Decode(&pokemon_stock)
	if err != nil {
		return data.StockPokemon{}
	}
	return pokemon_stock
}

func Get_all_pokemon() []data.StockPokemonView {
	stock_collection := MI.DB.Collection("stocks")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var pokemon_stocks []data.StockPokemonView
	cursor, err := stock_collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	if err != nil {
		return []data.StockPokemonView{}
	}
	for cursor.Next(ctx) {
		var StockPokemon data.StockPokemon
		var MinifiedPokemon data.MinifiedPokemon
		cursor.Decode(&StockPokemon)
		req.Do(fmt.Sprintf("pokemon/%s", StockPokemon.Name), &MinifiedPokemon)
		poke := &data.StockPokemonView{
			Pokemon:      MinifiedPokemon,
			StockPokemon: StockPokemon,
		}

		pokemon_stocks = append(pokemon_stocks, *poke)

	}

	return pokemon_stocks
}

func Destock_pokemon(name string, new_count int) *mongo.UpdateResult {
	stock_collection := MI.DB.Collection("stocks")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//var pokemon_stock data.StockPokemon
	filter := bson.M{"name": name}
	pof, object := stock_collection.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": bson.M{"count": new_count}},
	)
	fmt.Println(object)
	return pof
}
