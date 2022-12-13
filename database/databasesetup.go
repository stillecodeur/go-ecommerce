package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func GetDbClient() *mongo.Client {
	if mongoClient == nil {
		mongoClient = getMongoClient()
	}
	return mongoClient
}

func getMongoClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27018"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println("Successfully connected")
	return client
}

func UserCollection() *mongo.Collection {
	var collection *mongo.Collection = mongoClient.Database("Ecommerce").Collection("users")
	return collection
}

func ProductCollection() *mongo.Collection {
	var collection *mongo.Collection = mongoClient.Database("Ecommerce").Collection("products")
	return collection
}
