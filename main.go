package main

import (
	//"os"
	//"io/ioutil"
	"log"
	//"net/http"

	"fmt"
	"os"

	"github.com/joho/godotenv"
	"celebot/telegram/api"
)


func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	fmt.Printf("start polling")
	token := os.Getenv("BOTTOKEN")
	telegram.StartPolling(token)
}
