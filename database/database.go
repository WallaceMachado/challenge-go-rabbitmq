package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Db() *mongo.Client {

	uri := "mongodb://root:root@localhost:27017"
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)

	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
