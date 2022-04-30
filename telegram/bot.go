package telegram


type telegramBot struct {
	TOKEN string
	handlers map[string]EventHandler
}

func NewBot(token string) telegramBot {
	return telegramBot{TOKEN: token, handlers: map[string]EventHandler{}}
}
