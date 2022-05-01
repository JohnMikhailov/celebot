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

type handlersRegistry struct {
	handlers map[string]EventHandler
}

func newHandlersRegistry() handlersRegistry {
	return handlersRegistry{handlers: map[string]EventHandler{}}
}

func (registry handlersRegistry) addEventHandler(textCommand string, handler EventHandler) {
	registry.handlers[textCommand] = handler
}

func (registry handlersRegistry) handlerExists(commandName string) bool {
	if _, ok := registry.handlers[commandName]; ok {
		return true
	}
	return false
}

func (registry handlersRegistry) getTextHandlerByCommand(commandName string) EventHandler {
	if !registry.handlerExists(commandName) {
		return nil
	}
	val, _ := registry.handlers[commandName]
	return val
}
