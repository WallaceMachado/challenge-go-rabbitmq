package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wallacemachado/challenge-go-rabbitmq/router"
)

func main() {
	r := router.Generate()

	fmt.Println("Server is running!")
	log.Fatal(http.ListenAndServe(":3000", r))

}
