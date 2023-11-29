package database

import (
	"context"
	"fmt"
	"log"
	"scraper/imdb/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, databaseName string, collectionName string) *mongo.Collection {
	collection := client.Database(databaseName).Collection(collectionName)
	println(collection)
	return collection
}

// insert document to collections
func InsertCollection(client *mongo.Client, databaseName string, collectionName string, Shows []models.Show) {

	collection := client.Database(databaseName).Collection(collectionName)
	var interfaceArray []interface{}
	for _, show := range Shows {
		interfaceArray = append(interfaceArray, show)
	}
	insertManyResult, err := collection.InsertMany(context.TODO(), interfaceArray)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}
