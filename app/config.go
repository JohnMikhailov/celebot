package app

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type config struct {
	BOTTOKEN string
}

var config_ config

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config_.BOTTOKEN = os.Getenv("BOTTOKEN")
}

func GetConfig() *config {
	return &config_
}
