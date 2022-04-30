package telegram

import (
	"fmt"
	"strings"
)

type commandHandlers struct {
	textCommandHandlers map[string]TextCommandHandler
}

var registeredCommandHandlers = commandHandlers{textCommandHandlers: map[string]TextCommandHandler{}}

type TextCommandHandler interface {
	HandleTextCommand(params map[string]string, message Message)
}

func (handlers commandHandlers) addTextHandler(textCommand string, handler TextCommandHandler) {
	handlers.textCommandHandlers[textCommand] = handler
}

func AddTextHandler(textCommand string, handler TextCommandHandler) {
	registeredCommandHandlers.addTextHandler(textCommand, handler)
}

func (handlers commandHandlers) handlerExists(commandName string) bool {
	if _, ok := handlers.textCommandHandlers[commandName]; ok {
		return true	
	}

	return false
}

func (handlers commandHandlers) getTextHandlerByCommand(commandName string) TextCommandHandler {
	if !handlers.handlerExists(commandName) {
		return nil
	}
	val, _ := handlers.textCommandHandlers[commandName]
	return val
}

func ProcessMessage(message Message) {
	command, params := prepareTextCommand(message.Text)

	if !registeredCommandHandlers.handlerExists(command) {
		fmt.Println("Command handler not registered! Skiping message")
		return
	}

	handler := registeredCommandHandlers.getTextHandlerByCommand(command)
	handler.HandleTextCommand(params, message)
}

func prepareTextCommand(textCommand string) (string, map[string]string) {
	// command syntax: command param1=value1 param2=value2
	fmt.Println("raw message text:", textCommand)
	trancatedCommand := strings.Fields(textCommand)
	preparedCommand := trancatedCommand[0]

	var preparedParams = map[string]string{}

	params := trancatedCommand[1:]

	for _, param := range params {
		splitedParam := strings.Split(param, "=")
		if len(splitedParam) > 1 {
			paramName := splitedParam[0]
			paramValue := splitedParam[1]
			preparedParams[paramName] = paramValue
		}
	}

	fmt.Println("prepared params:", preparedParams)

	return preparedCommand, preparedParams
}
