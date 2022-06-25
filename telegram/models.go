package telegram

import (
	"reflect"
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

type forceReply struct {
	ForceReply bool `json:"force_reply"`
	InputFieldPlaceHolder string `json:"input_field_placeholder"`
	Selective bool `json:"selective"`
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
}

func (m *message) IsReply() bool {
	return reflect.ValueOf(m).Elem().FieldByName("ReplyToMessage") != reflect.Value{}
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

type KeyboardButton struct {
	// full description https://core.telegram.org/bots/api#keyboardbutton
	Text string `json:"text"`
}

type ReplyKeyboardMarkup struct {
	Keyboard [][]KeyboardButton `json:"keyboard"`
	OneTimeKeyboard bool `json:"one_time_keyboard"`
	Selective bool `json:"selective"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective bool `json:"selective"`
}

type Chat struct {
	ID int `json:"id"`
	Type string `json:"type"`
	Title string `json:"title"`
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}
