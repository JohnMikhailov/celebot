package telegram

import (
	"io/ioutil"
	"bytes"
	"log"
	"net/http"
	"strconv"

	"encoding/json"
)

type apiClient struct {
	urlHead string
	token string
}

func NewApiClient(token string) apiClient {
	return apiClient{token: token, urlHead: "https://api.telegram.org/bot"}
}

func (client apiClient) SendMessage(chatId, text string) *message {
	// add url query schema https://core.telegram.org/bots/api#sendmessage
	// TODO use url module
	// TODO add user-agent header
	res := message {}
	errorRes := errorResponse{}
	url := client.urlHead + client.token + "/sendMessage"

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
		log.Println(errorRes.Description)
	}

	return &res
}

func (client apiClient) getUpdates(updatesOffset int) *updateResponse {
	// add url query schema https://core.telegram.org/bots/api#getupdates
	// TODO use url module
	// TODO add user-agent header
	url := client.urlHead +
	client.token + "/" +
			"getUpdates" +
			"?timeout=10" +
			"&offset=" + strconv.Itoa(updatesOffset)

	res := updateResponse{}

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
