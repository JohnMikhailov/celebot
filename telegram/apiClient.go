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

	retriesCount int
	delayBetweenRetriesSec int
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
		retriesCount: 3,
		delayBetweenRetriesSec: 3,
	}
}

func (tc *telegramClient) send(request *http.Request) (*http.Response, error) {
	response, err := tc.httpClient.Do(request)

	if err != nil {
		log.Println("HTTP request failed", err.Error())
		return nil, err
	}

	return response, nil
}

func (tc *telegramClient) prepareRequestBody(requestBody *requestBodyType) (io.Reader, error) {
	if requestBody == nil {
		return nil, nil
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Println("Failed to marshal request body " + err.Error())
		return nil, err
	}

	return bytes.NewBuffer(jsonData), nil
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

func (tc *telegramClient) prepareRequest(method, urlTail string, requestBody *requestBodyType, queryParams *requestQueryParamsType) (*http.Request, error) {
	body, err := tc.prepareRequestBody(requestBody)

	if err != nil {
		return nil, err
	}

	url := tc.baseUrl + "/" + urlTail
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Failed to create request " + err.Error())
	}
	req.Header.Add("Content-Type", "application/json")

	tc.prepareQueryParams(queryParams, req)

	return req, nil
}

func (tc *telegramClient) getBodyBytes(response *http.Response) ([]byte, error) {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Failed to parse response body " + err.Error())
		return nil, err
	}

	return body, nil
}

func (tc *telegramClient) getTimeoutFromBody(body []byte) int {
	model := tooManyRequestsResponse{}
	json.Unmarshal(body, &model)
	return model.Parameters.RetryAfter
}

func (tc *telegramClient) wait(timeout int) error {
	time.Sleep(time.Duration(timeout) * time.Second)
	return nil
}

func (tc *telegramClient) sendRequest(method, urlTail string, body *requestBodyType, queryParams *requestQueryParamsType, responseModel interface{}) error {
	request, err := tc.prepareRequest(method, urlTail, body, queryParams)

	if err != nil {
		return err
	}

	for i := 1; i <= tc.retriesCount; i ++ {
		timeout := 0

		response, err := tc.send(request)

		if err != nil {
			return err
		}

		body, err := tc.getBodyBytes(response)

		if err != nil {
			return err
		}

		if response.StatusCode == http.StatusOK {
			json.Unmarshal(body, responseModel)
			return nil
		}

		log.Println("Bad status code:", response.StatusCode, "body:", string(body))

		timeout = tc.delayBetweenRetriesSec
		if response.StatusCode == http.StatusTooManyRequests {
			timeout = tc.getTimeoutFromBody(body)
		}

		log.Println("Attempt", i, "to call", urlTail, "failed with:", err.Error(), "next attempt in", timeout, "seconds")
		tc.wait(timeout)
	}

	log.Println("Maximum retries attempts exceeded for endpoint:", urlTail)

	return nil
}
