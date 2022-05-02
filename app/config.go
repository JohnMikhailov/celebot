package app

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type config struct {
	BOTTOKEN string

	DBUSERNAME string
    DBPASSWORD string
    DBHOST string
    DBSCHEMA string
}

var config_ config

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config_.BOTTOKEN = os.Getenv("BOTTOKEN")
	config_.DBUSERNAME = os.Getenv("DBUSERNAME")
	config_.DBPASSWORD = os.Getenv("DBPASSWORD")
	config_.DBHOST = os.Getenv("DBHOST")
	config_.DBSCHEMA = os.Getenv("DBSCHEMA")
}

func GetConfig() *config {
	return &config_
}
