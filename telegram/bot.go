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
