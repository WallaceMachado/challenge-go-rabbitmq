package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var (
	Port = ""

	DbUsername   = ""
	DbPassword   = ""
	DbName       = ""
	DbCollection = ""
	DbHost       = ""
	DbPort       = ""

	RabbitmqUser           = ""
	RabbitmqPass           = ""
	RabbitmqVhost          = ""
	RabbitmqHost           = ""
	RabbitmqPort           = ""
	RabbitmqExchangePerson = ""
)

func Init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}

	Port = os.Getenv("PORT")

	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	DbCollection = os.Getenv("DB_COLLECTION")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")

	RabbitmqUser = os.Getenv("RABBITMQ_DEFAULT_USER")
	RabbitmqPass = os.Getenv("RABBITMQ_DEFAULT_PASS")
	RabbitmqVhost = os.Getenv("RABBITMQ_DEFAULT_VHOST")
	RabbitmqHost = os.Getenv("RABBITMQ_DEFAULT_HOST")
	RabbitmqPort = os.Getenv("RABBITMQ_DEFAULT_PORT")
	RabbitmqExchangePerson = os.Getenv("RABBITMQ_EXCHANGE_PERSON")

}
