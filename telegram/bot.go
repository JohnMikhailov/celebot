package telegram

type telegramBot struct {
	TOKEN            string
	client           APICaller
	handlersRegistry handlersRegistry
}

func NewBot(token string) telegramBot {
	return telegramBot{
		TOKEN:            token,
		handlersRegistry: newHandlersRegistry(),
		client:           NewClient(token),
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

func (bot telegramBot) GetName() string {
	me, err := bot.client.GetMe()
	botname := "celebratorbot"
	if err == nil {
		botname = me.Username
	}

	return botname
}
