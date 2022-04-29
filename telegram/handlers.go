package telegram


import "strings"


var textCommandHandlers map[string]TextCommandHandler


type TextCommandHandler interface {
	Handle(params map[string]string, message Message)
}

func AddTextHandler(textCommand string, handler TextCommandHandler) {
	textCommandHandlers[textCommand] = handler
}

func HandleTextCommand(message Message) {
	command, params := prepareTextCommand(message.Text)
	handler, ok := textCommandHandlers[command]
	if ok {
		handler.Handle(params, message)
	}
}

func prepareTextCommand(textCommand string) (string, map[string]string) {
	// command syntax: command param1=value1 param2=value2
	trancatedCommand := strings.Fields(textCommand)
	preparedCommand := trancatedCommand[0]
	params := trancatedCommand[1:]

	var preparedParams map[string]string
	for _, param := range params {
		splitedParam := strings.Split(param, "=")
		paramName := splitedParam[0]
		paramValue := splitedParam[1]

		preparedParams[paramName] = paramValue
	}

	return preparedCommand, preparedParams
}
