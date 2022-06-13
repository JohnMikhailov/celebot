package app

import (
	"log"
	"os"
	"strconv"
	"time"
)

type config struct {
	BOTTOKEN string
	LONGPOLLING_WORKERS int
	DEFAULT_DELAY_BETWEEN_CHECKS_SEC time.Duration
}

var config_ config

func init() {
	config_.BOTTOKEN = os.Getenv("BOTTOKEN")
	longPolingWorkers, err := strconv.Atoi(os.Getenv("LONGPOLLING_WORKERS"))
	if err != nil {
		log.Fatalf("Parse var error for: LONGPOLLING_WORKERS")
	}
	config_.LONGPOLLING_WORKERS = longPolingWorkers
	seconds, err := strconv.Atoi(os.Getenv("DEFAULT_DELAY_BETWEEN_CHECKS_SEC"))
	if err != nil {
		log.Fatalf("Parse var error for: DEFAULT_DELAY_BETWEEN_CHECKS_SEC")
	}
	config_.DEFAULT_DELAY_BETWEEN_CHECKS_SEC = time.Duration(seconds) * time.Second
}

func GetConfig() *config {
	return &config_
}
