package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wallacemachado/challenge-go-rabbitmq/src/config"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/router"
)

func main() {
	config.Init()
	r := router.Generate()

	fmt.Println("Server is running!")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r))

}
