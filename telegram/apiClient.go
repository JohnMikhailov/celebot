package telegram

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"io"

	"encoding/json"
)

type requestBodyType map[string]interface{}
type requestQueryParamsType map[string]string

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

func (tc *telegramClient) prepareRequestBody(requestBody *requestBodyType) io.Reader {
	if requestBody == nil {
		return nil
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Failed to marshal request body " + err.Error())
	}

	return bytes.NewBuffer(jsonData)
}

func (tc *telegramClient) prepareQueryParams(queryParams *requestQueryParamsType, requset *http.Request) error {
	if queryParams == nil {
		return nil
	}

	query := requset.URL.Query()

	for key, value := range *queryParams {
		query.Add(key, value)
	}

	requset.URL.RawQuery = query.Encode()

	return nil
}

func (tc *telegramClient) prepareRequest(method, urlTail string, requestBody *requestBodyType, queryParams *requestQueryParamsType) *http.Request {
	body := tc.prepareRequestBody(requestBody)

	url := tc.baseUrl + "/" + urlTail
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Failed to create request " + err.Error())
	}
	req.Header.Add("Content-Type", "application/json")

	tc.prepareQueryParams(queryParams, req)

	return req
}

func (tc *telegramClient) sendRequest(method, urlTail string, body *requestBodyType, queryParams *requestQueryParamsType) []byte {
	request := tc.prepareRequest(method, urlTail, body, queryParams)
	return tc.send(request)
}
