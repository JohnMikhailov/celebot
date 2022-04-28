package app

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)


type Config struct {
	BOTTOKEN string
}

func Get(key string) string {
	return os.Getenv(key)
}

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
