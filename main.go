package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wallacemachado/challenge-go-rabbitmq/config"
	"github.com/wallacemachado/challenge-go-rabbitmq/database"
	"github.com/wallacemachado/challenge-go-rabbitmq/router"
)

func main() {
	config.Loader()
	r := router.Generate()
	database.Db().Database("challenge-go-rabbitmq-db").Collection("person")
	fmt.Println("Server is running!")
	log.Fatal(http.ListenAndServe(":3001", r))

}
