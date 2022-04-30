package main

import (
	//"os"
	//"io/ioutil"
	//"net/http"
	"github.com/meehighlov/celebot/telegram"
	"github.com/meehighlov/celebot/app"
)


type StartCommand struct {}

func (handler StartCommand) HandleTextCommand(params map[string]string, message telegram.Message) {
	telegram.SendMessage(
		app.GetConfig().BOTTOKEN,
		message.GetChatIdStr(),
		"Hello, i'm celebot! Tell me about your friends birthdays and i will remind you about it ;)",
	)
}

func main() {
	token := app.GetConfig().BOTTOKEN

	telegram.AddTextHandler("start", StartCommand{})
	telegram.StartPolling(token)
}
