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

	RabbitmqUser           = ""
	RabbitmqPass           = ""
	RabbitmqVhost          = ""
	RabbitmqHost           = ""
	RabbitmqPort           = ""
	RabbitmqExchangePerson = ""
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

	RabbitmqUser = os.Getenv("RABBITMQ_DEFAULT_USER")
	RabbitmqPass = os.Getenv("RABBITMQ_DEFAULT_PASS")
	RabbitmqVhost = os.Getenv("RABBITMQ_DEFAULT_VHOST")
	RabbitmqHost = os.Getenv("RABBITMQ_DEFAULT_HOST")
	RabbitmqPort = os.Getenv("RABBITMQ_DEFAULT_PORT")
	RabbitmqExchangePerson = os.Getenv("RABBITMQ_EXCHANGE_PERSON")

}
