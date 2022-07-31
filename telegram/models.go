package telegram

import (
	"reflect"
	"strconv"
	"strings"
)

type chat struct {
	// full description https://core.telegram.org/bots/api#chat
	//Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Id int `json:"id"`
	Type string `json:"type"`
	Title string `json:"title"`
	Username string `json:"username"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

type user struct {
	// full description https://core.telegram.org/bots/api#user
	Id int  `json:"id"`
	IsBot bool  `json:"is_bot"`
	FirstName string  `json:"first_name"`
	LastName string  `json:"last_name"`
	Username string  `json:"username"`
}

type chatMember struct {
	// full description https://core.telegram.org/bots/api#chatmemberowner
	Status string `json:"status"`
	User user `json:"user"`
}

type chatMemberResponse struct {
	Ok bool `json:"ok"`
	Result []chatMember `json:"result"`
}

type singleChatMemberResponse struct {
	Ok bool `json:"ok"`
	Result chatMember `json:"result"`
}

type replyToMessage struct {
	MessageId int  `json:"message_id"`
	From user  `json:"from"`
	SenderChat chat  `json:"sender_chat"`
	Chat chat `json:"chat"`
	Text string  `json:"text"`
}

type message struct {
	MessageId int  `json:"message_id"`
	From user  `json:"from"`
	SenderChat chat  `json:"sender_chat"`
	Chat chat `json:"chat"`
	Text string  `json:"text"`
	ReplyToMessage replyToMessage `json:"reply_to_message"`
	NewChatMembers []user `json:"new_chat_members"`
	LeftChatMember user `json:"left_chat_member"`
}

func (m *message) IsReply() bool {
	return reflect.ValueOf(m).Elem().FieldByName("ReplyToMessage") != reflect.Value{}
}

func (m *message) HasLeftChatMember() bool {
	return reflect.ValueOf(m).Elem().FieldByName("LeftChatMember") != reflect.Value{}
}

type update struct {
	UpdateId int `json:"update_id"`
	Message message `json:"message"`
}

func (update *update) isFromGroup() bool {
	return update.Message.Chat.Type == "group" // todo use enum!
}

func (update *update) isFromPrivateChat() bool {
	return update.Message.Chat.Type == "private" // todo use enum!
}

type getMeReponse struct {
	Ok bool `json:"ok"`
	Result user `json:"result"`
}

type updateResponse struct {
	Ok bool `json:"ok"`
	Result []update `json:"result"`
}

func (update *updateResponse) GetLastUpdateId() int {
	return update.Result[len(update.Result) - 1].UpdateId
}

func (message message) GetChatIdStr() string {
	return strconv.Itoa(message.Chat.Id)
}

func (message message) GetSenderChatIdStr() string {
	return strconv.Itoa(message.SenderChat.Id)
}

func (message message) getCommand() string {
	parts := strings.Fields(message.Text)
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

type retryAfter struct {
	RetryAfter int `json:"retry_after"`
}

type tooManyRequestsResponse struct {
	Ok bool `json:"ok"`
	ErrorCode int `json:"error_code"`
	Description string `json:"description"`
	Parameters retryAfter `json:"parameters"`
}
