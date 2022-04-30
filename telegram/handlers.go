package telegram

import (
	"fmt"
	"strings"
)

type TextCommandHandler interface {
	HandleTextCommand(bot telegramBot, params map[string]string, message Message)
}

type MessageProcessor interface {
	ProcessMessage(message Message)
}

func (bot telegramBot) AddTextHandler(textCommand string, handler TextCommandHandler) {
	if bot.textCommandHandlers == nil {
		bot.textCommandHandlers = make(map[string]TextCommandHandler)
		fmt.Println("handlers map is empty, created")
	}
	bot.textCommandHandlers[textCommand] = handler
	fmt.Println(bot.textCommandHandlers)
}

func (bot telegramBot) handlerExists(commandName string) bool {
	if bot.textCommandHandlers == nil {
		fmt.Println("pizdeeeeeeeec")
	}
	if _, ok := bot.textCommandHandlers[commandName]; ok {
		return true
	}
	return false
}

func (bot telegramBot) getTextHandlerByCommand(commandName string) TextCommandHandler {
	if !bot.handlerExists(commandName) {
		return nil
	}
	val, _ := bot.textCommandHandlers[commandName]
	return val
}

func (bot telegramBot) processMessage(message Message) {
	command, params := prepareTextCommand(message.Text)

	if !bot.handlerExists(command) {
		fmt.Println("Command handler not registered! Skiping message")
		return
	}

	handler := bot.getTextHandlerByCommand(command)
	handler.HandleTextCommand(bot, params, message)
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
