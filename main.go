package main

import (
	"os"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"telegram/api"
)


func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {

}
