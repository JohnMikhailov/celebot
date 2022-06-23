package telegram


import "strings"


type Bundle interface {
	Bot() *telegramBot
	Message() *message
	SendMessage(text, chatId string, needForceReply bool) *message
	SendMessageWithKeyboard(chatId string, keyboard ReplyKeyboardMarkup) *message
	Args() []string
}

type bundle struct {
	bot *telegramBot
	message *message
}

func newBundle(bot *telegramBot, message *message) *bundle {
	return &bundle{bot: bot, message: message}
}

func (b *bundle) Message() *message {
	return b.message
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

func (b *bundle) SendMessage(text, chatId string, needForceReply bool) *message {
	return b.Bot().client.SendMessage(text, chatId, needForceReply)
}

func (b *bundle) SendMessageWithKeyboard(chatId string, keyboard ReplyKeyboardMarkup) *message {
	return b.Bot().client.SendKeyboard(chatId, keyboard)
}
