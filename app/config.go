package app

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)


type Config struct {
	BOTTOKEN string
}

var config Config

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config.BOTTOKEN = os.Getenv("BOTTOKEN")
}

func GetConfig() *Config {
	return &config
}
