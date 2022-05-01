package telegram


type EventHandler interface {
	OnEvent(e Event)
}

type Event struct {
	bot telegramBot
	Message Message
}

func (e Event) SendMessage(text, chatId string) *Message {
	return e.bot.client.sendMessage(text, chatId)
}
