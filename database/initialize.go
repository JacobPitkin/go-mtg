package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"jacobpitkin.com/go-mtg/cards"
)

// func init() {
// 	_, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB on init")
// }

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("mtg").Collection("cards")
	cmc3 := bson.M{"cmc": 3}

	var result cards.CardList
	cursor, err := coll.Find(context.TODO(), cmc3)
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB on Connect")
	fmt.Println(result)
}
