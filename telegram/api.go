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

func (client apiClient) SendMessage(chatId, text string, needForceReply bool) *message {
	// add url query schema https://core.telegram.org/bots/api#sendmessage
	// TODO use url module
	// TODO add user-agent header
	// TODO use model for body
	res := message {}
	errorRes := errorResponse{}
	url := client.urlHead + client.token + "/sendMessage"

	body := map[string]string{"chat_id": chatId, "text": text}
	if needForceReply {
		force_reply_json, err := json.Marshal(map[string]bool{"force_reply": true, "selective": true})
		if err != nil {
			log.Fatal(err)
		}
		body["reply_markup"] = string(force_reply_json)
	}
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
	errorRes := errorResponse{}

	resp, err := http.Get(url)

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

func (client apiClient) SendKeyboard(chatId string, keyboard ReplyKeyboardMarkup) *message {
	// add url query schema https://core.telegram.org/bots/api#sendmessage
	// TODO use url module
	// TODO add user-agent header
	// TODO use model for body
	res := message {}
	errorRes := errorResponse{}
	url := client.urlHead + client.token + "/sendMessage"

	body := map[string]string{"chat_id": chatId}
	keyboard_json, err := json.Marshal(keyboard)
	if err != nil {
		log.Fatal(err)
	}
	body["reply_markup"] = string(keyboard_json)

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

func (client apiClient) GetChatAdministrators(chatId string) (*[]chatMember, error) {
	// TODO use url module
	// TODO add user-agent header
	// TODO use model for body
	url := client.urlHead + client.token + "/getChatAdministrators"

	body := map[string]string{"chat_id": chatId}
	json_data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := chatMemberResponse{}
	errorRes := errorResponse{}

	bodyString := string(bodyBytes)
	json.Unmarshal([]byte(bodyString), &res)

	if resp.StatusCode == http.StatusOK {
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal([]byte(bodyString), &res)
	} else {
		json.Unmarshal([]byte(bodyString), &errorRes)
		log.Println(errorRes.Description)
	}

	return &res.Result, nil
}
