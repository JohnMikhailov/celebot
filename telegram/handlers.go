package telegram

import (
	"fmt"
)

type EventHandler interface {
	OnEvent(e Event)
}

type Event struct {
	bot telegramBot
	Message Message
}

func (e Event) SendMessage(text, chatId string) *Message {
	return e.bot.sendMessage(text, chatId)
}

func (bot telegramBot) AddEventHandler(textCommand string, handler EventHandler) {
	bot.handlers[textCommand] = handler
}

func (bot telegramBot) handlerExists(commandName string) bool {
	if _, ok := bot.handlers[commandName]; ok {
		return true
	}
	return false
}

func (bot telegramBot) getTextHandlerByCommand(commandName string) EventHandler {
	if !bot.handlerExists(commandName) {
		return nil
	}
	val, _ := bot.handlers[commandName]
	return val
}

func (bot telegramBot) processMessage(message Message) {
	command := message.getCommand()
	if !bot.handlerExists(command) {
		fmt.Println("Command handler not registered! Skiping message")
		return
	}

	handler := bot.getTextHandlerByCommand(command)
	event := Event{bot: bot, Message: message}
	handler.OnEvent(event)
}
