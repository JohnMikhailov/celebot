package app

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type config struct {
	BOTTOKEN string
	LONGPOLLING_WORKERS int
	DEFAULT_DELAY_BETWEEN_REMINDINGS_SEC time.Duration
}

var config_ config

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config_.BOTTOKEN = os.Getenv("BOTTOKEN")
	config_.LONGPOLLING_WORKERS, err = strconv.Atoi(os.Getenv("LONGPOLLING_WORKERS"))
	if err != nil {
		log.Fatalf("Parse .env error for: LONGPOLLING_WORKERS")
	}
	seconds, err := strconv.Atoi(os.Getenv("DEFAULT_DELAY_BETWEEN_REMINDINGS_SEC"))
	if err != nil {
		log.Fatalf("Parse .env error for: DEFAULT_DELAY_BETWEEN_REMINDINGS_SEC")
	}
	config_.DEFAULT_DELAY_BETWEEN_REMINDINGS_SEC = time.Duration(seconds) * time.Second
}

func GetConfig() *config {
	return &config_
}
