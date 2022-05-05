package mongoconnect

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

func CheckConnect() {

	/*
	   Connect to my cluster
	*/
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

	/*
	   List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

}

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
	for i := 0; i < 100; i++ {
		res, _ := generate(rand.Intn(500))
		res.Count = rand.Intn(7)
		res.Price = rand.Intn(150)
		//res2B, _ := json.Marshal(res.Pokemon)
		//fmt.Println(string(res2B))
		//b, _ := json.Marshal(res)
		stocks.InsertOne(ctx, res)
	}
}

func generate(id int) (result data.StockPokemon, err error) {
	err = req.Do(fmt.Sprintf("pokemon/%d", id), &result.Pokemon)
	return result, err
}
