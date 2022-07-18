package telegram

import (
	"encoding/json"
	"strconv"
)

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

	responseBytes := tc.sendRequest("POST", "sendMessage", &body, nil)

	json.Unmarshal(responseBytes, &res)

	return &res
}

func (tc telegramClient) GetUpdates(updatesOffset int) *updateResponse {
	res := updateResponse{}

	queryParams := requestQueryParamsType{
		"timeout": "10",
		"offset": strconv.Itoa(updatesOffset),
	}

	responseBytes := tc.sendRequest("GET", "getUpdates", nil, &queryParams)

	json.Unmarshal(responseBytes, &res)

	return &res
}

func (tc telegramClient) GetMe() (*user, error) {
	res := getMeReponse{}

	responseBytes := tc.sendRequest("GET", "getMe", nil, nil)

	json.Unmarshal(responseBytes, &res)

	return &res.Result, nil
}

func (tc telegramClient) GetChatAdministrators(chatId string) (*[]chatMember, error) {
	res := chatMemberResponse{}
	body := requestBodyType{"chat_id": chatId}

	responseBytes := tc.sendRequest("GET", "getChatAdministrators", &body, nil)
	json.Unmarshal(responseBytes, &res)

	return &res.Result, nil
}

func (tc telegramClient) GetChatMember(chatId, userId string) (*singleChatMemberResponse, error) {
	res := singleChatMemberResponse{}
	body := requestBodyType{"chat_id": chatId, "user_id": userId}

	responseBytes := tc.sendRequest("GET", "getChatMember", &body, nil)
	json.Unmarshal(responseBytes, &res)

	return &res, nil
}
