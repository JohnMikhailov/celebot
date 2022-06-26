package telegram

import (
	"strconv"
	"strings"
)


type Bundle interface {
	Bot() *telegramBot
	Message() *message
	Update() *update
	SendMessage(text, chatId string, needForceReply bool) *message
	GetChatAdministrators(chatId string) (*[]chatMember, error)
	GetChatOwner(chatId string) (*chatMember, error)
	GetChatMember(chatId, userId int) (*chatMember, error)
	GetMe() (*user, error)
	Args() []string
}

type bundle struct {
	bot *telegramBot
	message *message
	update *update
}

func newBundle(bot *telegramBot, message *message, update *update) *bundle {
	return &bundle{bot: bot, message: message, update: update}
}

func (b *bundle) Message() *message {
	return b.message
}

func (b *bundle) Update() *update {
	return b.update
}

func (b *bundle) Bot() *telegramBot {
	return b.bot
}

func (b *bundle) Args() []string {
	args := strings.Fields(b.Message().Text)
	if len(args) > 1 {
		return args[1:]
	}
	return []string{}
}

func (b *bundle) GetMe() (*user, error) {
	return b.bot.client.GetMe()
}

func (b *bundle) SendMessage(text, chatId string, needForceReply bool) *message {
	return b.Bot().client.SendMessage(text, chatId, needForceReply)
}

func (b *bundle) GetChatAdministrators(chatId string) (*[]chatMember, error) {
	return b.Bot().client.GetChatAdministrators(chatId)
}

func (b *bundle) GetChatOwner(chatId string) (*chatMember, error) {
	admins, err := b.Bot().client.GetChatAdministrators(chatId)
	if err != nil {
		return nil, err
	}

	for _, admin := range *admins {
		if admin.Status == "creator" {
			return &admin, nil
		}
	}

	return nil, nil
}

func (b *bundle) GetChatMember(chatId, userId int) (*chatMember, error) {
	chatIdStr := strconv.Itoa(chatId)
	userIdStr := strconv.Itoa(userId)

	return b.bot.client.GetChatMember(chatIdStr, userIdStr)
}
