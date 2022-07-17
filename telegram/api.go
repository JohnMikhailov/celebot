package telegram

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"encoding/json"
)

type requestBodyType map[string]interface{}

type telegramClient struct {
	urlHead string
	token   string
	baseUrl string

	httpClient *http.Client
}

type APICaller interface {
	SendMessage(chatId, text string, needForceReply bool) *message
	GetUpdates(updatesOffset int) *updateResponse
	GetChatAdministrators(chatId string) (*[]chatMember, error)
	GetMe() (*user, error)
	GetChatMember(chatId, userId string) (*singleChatMemberResponse, error)
}

func NewClient(token string) APICaller {
	httpClient := &http.Client{Timeout: 20 * time.Second}
	urlHead := "https://api.telegram.org/bot"
	return telegramClient{
		token: token,
		urlHead: urlHead,
		baseUrl: urlHead + token,
		httpClient: httpClient,
	}
}

func (client telegramClient) GetUpdates(updatesOffset int) *updateResponse {
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

func (client telegramClient) GetChatAdministrators(chatId string) (*[]chatMember, error) {
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

func (client telegramClient) GetChatMember(chatId, userId string) (*singleChatMemberResponse, error) {
	// TODO use url module
	// TODO add user-agent header
	// TODO use model for body
	url := client.urlHead + client.token + "/getChatMember"

	body := map[string]string{"chat_id": chatId, "user_id": userId}
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

	res := singleChatMemberResponse{}
	errorRes := errorResponse{}

	bodyString := string(bodyBytes)

	if resp.StatusCode == http.StatusOK {
		if err != nil {

			return nil, err
		}
		json.Unmarshal([]byte(bodyString), &res)
	} else {
		json.Unmarshal([]byte(bodyString), &errorRes)
		log.Println(errorRes.Description)
	}

	return &res, nil
}

func (client telegramClient) GetMe() (*user, error) {
	// TODO use url module
	// TODO add user-agent header
	// TODO use model for body
	url := client.urlHead + client.token + "/getMe"

	resp, err := http.Get(url)

	if err != nil {
		log.Println("Error getting info about bot: " + err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error forming response from : " + err.Error())
		return nil, err
	}

	res := getMeReponse{}
	errorRes := errorResponse{}

	bodyString := string(bodyBytes)

	if resp.StatusCode == http.StatusOK {
		json.Unmarshal([]byte(bodyString), &res)
	} else {
		json.Unmarshal([]byte(bodyString), &errorRes)
		log.Println(errorRes.Description)
	}

	return &res.Result, nil
}

func (tc *telegramClient) send(request *http.Request) []byte {
	response, err := tc.httpClient.Do(request)

	if err != nil {
		log.Fatalf("HTTP request failed " + err.Error())
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to parse response body " + err.Error())
	}

	if response.StatusCode != http.StatusOK {
		log.Println("Bad status code: ", response.StatusCode, " body: ",string(body))
		return []byte{}
	}

	return body
}

func (tc *telegramClient) prepareRequest(method, urlTail string, requestBody *requestBodyType) *http.Request {
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Failed to marshal request body for " + method + " " + err.Error())
	}
	log.Println(bytes.NewBuffer(jsonData))

	url := tc.baseUrl + "/" + urlTail
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to create request " + err.Error())
	}
	req.Header.Add("ContentType", "application/json")

	log.Println("prepared request: ", req.Body)

	return req
}

func (tc *telegramClient) sendRequest(method, urlTail string, body *requestBodyType) []byte {
	request := tc.prepareRequest(method, urlTail, body)
	return tc.send(request)
}

func (tc telegramClient) SendMessage(chatId, text string, needForceReply bool) *message {
	res := message{}

	body := requestBodyType{
		"chat_id": chatId,
		"text": text,
		"reply_markup": requestBodyType{
			"force_reply": needForceReply,
			"selective": needForceReply,
		},
	}

	responseBytes := tc.sendRequest("POST", "sendMessage", &body)

	json.Unmarshal(responseBytes, &res)

	return &res
}
