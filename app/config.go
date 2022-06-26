package app

import (
	"log"
	"os"
	"strconv"
	"time"
)

type empty int

type config struct {
	BOTTOKEN_CELEBOT                 string
	LONGPOLLING_WORKERS              int
	DEFAULT_DELAY_BETWEEN_CHECKS_SEC time.Duration
	ALLOWED_USERS                    map[string]empty
	BD_NOTIFICATION_HOUR_MOSCOW_TZ   int
}

var config_ config

func init() {
	config_.BOTTOKEN_CELEBOT = os.Getenv("BOTTOKEN_CELEBOT")
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
	config_.BD_NOTIFICATION_HOUR_MOSCOW_TZ, err = strconv.Atoi(os.Getenv("BD_NOTIFICATION_HOUR_MOSCOW_TZ"))
	if err != nil {
		log.Fatalf("Parse var error for: BD_NOTIFICATION_HOUR_MOSCOW_TZ")
	}
}

func GetConfig() *config {
	return &config_
}

func (c *config) IsUsernameExist(username string) bool {
	if _, ok := c.ALLOWED_USERS[username]; ok {
		return true
	}
	return false
}
