package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println("failed to ping mongo")
		return nil
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {

}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {

}
