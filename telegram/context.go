package telegram


type Bundle interface {
	Bot() *telegramBot
	Message() *message
	SendMessage(text, chatId string) *message
	SendMessageWithKeyboard(text, chatId string, keyboard ReplyKeyboardMarkup) *message
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

func (b *bundle) SendMessage(text, chatId string) *message {
	return b.Bot().client.SendMessage(text, chatId, true)
}

func (b *bundle) SendMessageWithKeyboard(text, chatId string, keyboard ReplyKeyboardMarkup) *message {
	return b.Bot().client.SendKeyboard(text, chatId, keyboard)
}
