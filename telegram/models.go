package telegram


import (
	"strconv"
	"strings"
)

type chat struct {
	// full description https://core.telegram.org/bots/api#chat

	Id int `json:"id"`

	//Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string  `json:"type"`

	Title string  `json:"title"`
	Username string  `json:"username"`
}

type user struct {
	// full description https://core.telegram.org/bots/api#user

	Id int  `json:"id"`
	IsBot bool  `json:"is_bot"`
	FirstName string  `json:"first_name"`
	LastName string  `json:"last_name"`
	Username string  `json:"username"`
}

type message struct {
	Message_id int  `json:"message_id"`
	From user  `json:"from"`
	SenderChat chat  `json:"sender_chat"`
	Chat chat `json:"chat"`
	Text string  `json:"text"`
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

type updateResponse struct {
	Ok bool `json:"ok"`
	Result []update `json:"result"`
}

func (update *updateResponse) GetLastUpdateId() int {
	return update.Result[len(update.Result) - 1].UpdateId
}

type responseParameters struct {
	MigrateToChatId int `json:"migrate_to_chat_id"`
	RetryAfter int `json:"retry_after"`
}

type errorResponse struct {
	Ok bool `json:"ok"`
	Description string `json:"description"`
	Parameters responseParameters  `json:"parameters"`
}

func (message message) GetChatIdStr() string {
	return strconv.Itoa(message.Chat.Id)
}

func (message message) getCommand() string {
	parts := strings.Fields(message.Text)
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}
