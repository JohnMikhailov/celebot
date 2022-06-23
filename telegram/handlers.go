package telegram

// TODO use generics

type handlerType func(bundle Bundle) error

type handlersRegistry struct {
	handlers    map[string]handlerType
	replyHandlers map[string]handlerType
}

func newHandlersRegistry() handlersRegistry {
	return handlersRegistry{handlers: map[string]handlerType{}, replyHandlers: map[string]handlerType{}}
}

func (registry *handlersRegistry) handlerExists(commandName string) bool {
	if _, ok := registry.handlers[commandName]; ok {
		return true
	}
	return false
}

func (registry *handlersRegistry) getHandlerByCommand(commandName string) handlerType {
	if !registry.handlerExists(commandName) {
		return nil
	}
	val := registry.handlers[commandName]
	return val
}

func (registry *handlersRegistry) replyHandlerExist(replyText string) bool {
	if _, ok := registry.replyHandlers[replyText]; ok {
		return true
	}
	return false
}

func (registry *handlersRegistry) getReplyHandlerByMessageText(replyText string) handlerType {
	if !registry.replyHandlerExist(replyText) {
		return nil
	}
	val := registry.replyHandlers[replyText]
	return val
}

func (registry *handlersRegistry) addHandler(commandName string, handler handlerType) {
	registry.handlers[commandName] = handler
}

func (registry *handlersRegistry) addReplyHandler(replyText string, handler handlerType) {
	registry.replyHandlers[replyText] = handler
}
