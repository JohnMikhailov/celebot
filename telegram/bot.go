package telegram

type telegramBot struct {
	TOKEN            string
	client           apiClient
	handlersRegistry handlersRegistry
}

func NewBot(token string) telegramBot {
	return telegramBot{
		TOKEN:            token,
		handlersRegistry: newHandlersRegistry(),
		client:           NewApiClient(token),
	}
}

func (bot telegramBot) AddHandler(textCommand string, handler handlerType) {
	bot.handlersRegistry.addHandler(textCommand, handler)
}

func (bot telegramBot) AddReplyHandler(replyText string, handler handlerType) {
	bot.handlersRegistry.addReplyHandler(replyText, handler)
}

func (bot telegramBot) SetDefaultHandler(handler handlerType) {
	bot.handlersRegistry.addDefaultHandler(handler)
}
