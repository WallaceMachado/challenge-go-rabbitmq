package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port = ""

	DbUsername   = ""
	DbPassword   = ""
	DbName       = ""
	DbCollection = ""
)

func Loader() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port = os.Getenv("PORT")

	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	DbCollection = os.Getenv("DB_COLLECTION")

}
