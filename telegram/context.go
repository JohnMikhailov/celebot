package telegram


type MessageHandler interface {
	Handle(c *Context)
}

type Context struct {
	bot telegramBot
	Message message
}

func (c *Context) SendMessage(text, chatId string, needForceReply bool) *message {
	return c.bot.client.SendMessage(text, chatId, needForceReply)
}

func (c *Context) SendMessageWithKeyboard(text, chatId string, keyboard ReplyKeyboardMarkup) *message {
	return c.bot.client.SendKeyboard(text, chatId, keyboard)
}
