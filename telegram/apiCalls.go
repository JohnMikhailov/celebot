package telegram

import (
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

	tc.sendRequest("POST", "sendMessage", &body, nil, res)

	return &res
}

func (tc telegramClient) GetUpdates(updatesOffset int) *updateResponse {
	res := updateResponse{}

	queryParams := requestQueryParamsType{
		"timeout": "10",
		"offset": strconv.Itoa(updatesOffset),
	}

	tc.sendRequest("GET", "getUpdates", nil, &queryParams, &res)

	return &res
}

func (tc telegramClient) GetMe() (*user, error) {
	res := getMeReponse{}
	tc.sendRequest("GET", "getMe", nil, nil, &res)
	return &res.Result, nil
}

func (tc telegramClient) GetChatAdministrators(chatId string) (*[]chatMember, error) {
	res := chatMemberResponse{}
	body := requestBodyType{"chat_id": chatId}

	tc.sendRequest("GET", "getChatAdministrators", &body, nil, &res)

	return &res.Result, nil
}

func (tc telegramClient) GetChatMember(chatId, userId string) (*singleChatMemberResponse, error) {
	res := singleChatMemberResponse{}
	body := requestBodyType{"chat_id": chatId, "user_id": userId}

	tc.sendRequest("GET", "getChatMember", &body, nil, &res)
	return &res, nil
}
