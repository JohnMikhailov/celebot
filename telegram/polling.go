package telegram


import "fmt"


func (bot telegramBot) StartPolling() {
	updatesOffset := -1
	fmt.Println("start polling")
	updatesQueue := make(chan update, 1000)

	go func() {
		update := <- updatesQueue
		bot.processMessage(update.Message)
	}()

	for {
		updates := bot.client.getUpdates(updatesOffset)
		if !updates.Ok {
			fmt.Println("getting updates failed")  // TODO log it
		}
		if len(updates.Result) > 0 {
			updatesOffset = updates.GetLastUpdateId() + 1

			for _, update := range updates.Result {

				updatesQueue <- update
			}
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
