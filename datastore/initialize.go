package datastore

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"jacobpitkin.com/go-mtg/cards"
)

var client *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoUri := os.Getenv("MONGO_URI")
	mongoClient, err := mongo.Connect(options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}

	client = mongoClient
	fmt.Println("Connected to MongoDB on init")
}

func Connect() {
	coll := client.Database("mtg").Collection("cards")
	filter := bson.M{"cmc": 3}

	var result cards.CardList
	startTime := time.Now().UnixMilli()
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	endTime := time.Now().UnixMilli()

	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(result)

	fmt.Printf("Number of CMC 3 results: %d\n", len(result))

	result = result.IsEligibleCommander()

	// fmt.Println(result)
	fmt.Printf("Number of CMC 3 eligible commanders: %d\n", len(result))

	fmt.Printf("Query time: %d ms\n", endTime-startTime)
}
