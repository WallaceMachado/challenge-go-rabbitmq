package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wallacemachado/challenge-go-rabbitmq/src/config"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/database"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/router"
)

func main() {
	config.Loader()
	r := router.Generate()
	database.Db().Database("challenge-go-rabbitmq-db").Collection("person")
	fmt.Println("Server is running!")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r))

}
