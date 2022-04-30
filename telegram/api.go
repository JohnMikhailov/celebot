package telegram

import (
	"io/ioutil"
	"bytes"
	"log"
	"net/http"
	"fmt"
	"strconv"

	"encoding/json"
)

const TelegramBotApiUrl = "https://api.telegram.org/bot"

var updatesClient = http.Client{}


func sendMessage(token, chatId, text string) *Message {
	res := Message {}
	errorRes := ErrorResponse{}
	url := TelegramBotApiUrl + token + "/sendMessage"

	body := map[string]string{"chat_id": chatId, "text": text}
	json_data, err := json.Marshal(body)

	if err != nil {
        log.Fatal(err)
    }

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if resp.StatusCode == http.StatusOK {
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal([]byte(bodyString), &res)
	} else {
		json.Unmarshal([]byte(bodyString), &errorRes)
		fmt.Println(errorRes.Description)
	}

	return &res
}

func SendMessage(token, chatId, text string) *Message {
	return sendMessage(token, chatId, text)
}

func getUpdates(token string, updatesOffset int) *UpdateResponse {
	url := TelegramBotApiUrl +
		    token + "/" +
			"getUpdates" +
			"?timeout=10" +
			"&limit=1" +
			"&offset=" + strconv.Itoa(updatesOffset)

	res := UpdateResponse{}

	resp, err := updatesClient.Get(url)

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
	updatesOffset := -1
	fmt.Println("start polling")
	// TODO: support another types of handlers
	for {
		updates := getUpdates(token, updatesOffset)
		if len(updates.Result) > 0 {
			updatesOffset = updates.Result[0].UpdateId + 1
		    message := updates.Result[0].Message
			ProcessMessage(message)
		} else {
			fmt.Println("no updates yet\n")
		}
	}
}
