package telegram

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
	"strconv"
	// "time"

	"encoding/json"
)

const TELEGRAMBOTURL = "https://api.telegram.org/bot"


func sendMessage(token, chatId, text string) *Message {
	res := Message {}
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


func getUpdates(token string, client *http.Client, updatesOffset int) *UpdateResponse {
	url := TELEGRAMBOTURL +
		    token + "/" +
			"getUpdates" +
			"?timeout=10" +
			"&limit=1" +
			"&offset=" + strconv.Itoa(updatesOffset)

	// url += "&" + strconv.Itoa(updatesOffset)

	res := UpdateResponse{}

	resp, err := client.Get(url)

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
	updatesClient := http.Client{}
	updatesOffset := -1
	for {
		updates := getUpdates(token, &updatesClient, updatesOffset)
		// chatId := strconv.Itoa(updates.Result[0].MessageInfo.SenderChat.Id)
		// text := "hello"
		// message := sendMessage(token, chatId, text)
		if len(updates.Result) > 0 {
			updatesOffset = updates.Result[0].UpdateId + 1
			fmt.Print(updates.Result[0].MessageInfo.Text)
		} else {
			fmt.Print("no updates yet\n")
		}
	}
}
