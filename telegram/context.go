package telegram


type MessageHandler interface {
	Handle(c *Context)
}

type Context struct {
	bot telegramBot
	Message message
}

func (c *Context) SendMessage(text, chatId string) *message {
	return c.bot.client.sendMessage(text, chatId)
}
