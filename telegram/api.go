package telegram

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
	"strconv"

	"encoding/json"
	"celebot/telegram/models"
)

const TELEGRAMBOTURL = "https://api.telegram.org/bot"


func sendMessage(token, chatId, text string) *telegram.Message {
	res := telegram.Message {}
	url := TELEGRAMBOTURL +
	       token + "/" +
		   "sendMessage" +
		   "?chat_id=" + chatId +
		   "&text=" + text

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		json.Unmarshal([]byte(bodyString), &res)
	}

	return &res
}


func getUpdates(token string) *telegram.UpdateResponse {
	url := TELEGRAMBOTURL +
		    token + "/" +
			"getUpdates" +
			"?timeout=5" +
			"&limit=1"

	res := telegram.UpdateResponse {}

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		json.Unmarshal([]byte(bodyString), &res)
	}

	return &res
}


func StartPolling(token string) {
	for {
		updates := getUpdates(token)
		chatId := strconv.Itoa(updates.Result[0].MessageInfo.SenderChat.Id)
		text := "hello"
		message := sendMessage(token, chatId, text)
		fmt.Print(message.Text)
	}
}
