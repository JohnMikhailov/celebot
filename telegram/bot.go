package telegram


type telegramBot struct {
	TOKEN string
	client apiClient
	handlersRegistry handlersRegistry
}

func NewBot(token string) telegramBot {
	return telegramBot{
		TOKEN: token,
		handlersRegistry: newHandlersRegistry(),
		client: NewApiClient(token),
	}
}

func (bot telegramBot) AddEventHandler(textCommand string, handler MessageHandler) {
	bot.handlersRegistry.addEventHandler(textCommand, handler)
}
