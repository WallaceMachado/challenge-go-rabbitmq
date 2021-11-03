package database

import (
	"context"
	"fmt"
	"log"

	"github.com/wallacemachado/challenge-go-rabbitmq/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func Db() *mongo.Client {

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)

	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db = client

	return client
}

func DbClose() {
	db.Disconnect(context.TODO())
}
