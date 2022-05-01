package telegram


type MessageHandler interface {
	Handle(c Context)
}

type Context struct {
	bot telegramBot
	Message Message
}

func (c Context) SendMessage(text, chatId string) *Message {
	return c.bot.client.sendMessage(text, chatId)
}
