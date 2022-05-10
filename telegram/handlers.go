package telegram


type handlersRegistry struct {
	handlers map[string]MessageHandler
}

func newHandlersRegistry() handlersRegistry {
	return handlersRegistry{handlers: map[string]MessageHandler{}}
}

func (registry *handlersRegistry) addEventHandler(textCommand string, handler MessageHandler) {
	registry.handlers[textCommand] = handler
}

func (registry *handlersRegistry) handlerExists(commandName string) bool {
	if _, ok := registry.handlers[commandName]; ok {
		return true
	}
	return false
}

func (registry *handlersRegistry) getTextHandlerByCommand(commandName string) MessageHandler {
	if !registry.handlerExists(commandName) {
		return nil
	}
	val := registry.handlers[commandName]
	return val
}
