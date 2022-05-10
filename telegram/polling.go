package telegram


import "fmt"


func (bot telegramBot) StartPolling() {
	updatesOffset := -1
	fmt.Println("start polling")
	// TODO: support another types of handlers
	for {
		updates := bot.client.getUpdates(updatesOffset)
		if len(updates.Result) > 0 {
			updatesOffset = updates.Result[0].UpdateId + 1
		    message := updates.Result[0].Message
			bot.processMessage(message)
		} else {
			fmt.Println("no updates yet")
		}
	}
}

func (bot telegramBot) processMessage(message message) {
	command := message.getCommand()
	if !bot.handlersRegistry.handlerExists(command) {
		fmt.Println("Command handler not registered! Skiping message")
		return
	}

	handler := bot.handlersRegistry.getTextHandlerByCommand(command)
	context := Context{bot: bot, Message: message}
	handler.Handle(&context)
}
