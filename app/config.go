package app

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type empty int

type config struct {
	BOTTOKEN_CELEBOT string
	LONGPOLLING_WORKERS int
	DEFAULT_DELAY_BETWEEN_CHECKS_SEC time.Duration
	ALLOWED_USERS map[string]empty
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
	config_.ALLOWED_USERS = make(map[string]empty)
	allowedUsers := os.Getenv("ALLOWED_USERS")
	allowedUsersList := strings.Split(allowedUsers, ",")
	for _, allowedUser := range allowedUsersList {
		config_.ALLOWED_USERS[allowedUser] = 1
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
