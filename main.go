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
	bot.SendMessage(
		message.GetChatIdStr(),
		"Hello, i'm celebot! Tell me about your friends birthdays and i will remind you about it ;)",
	)
}

func main() {
	token := app.GetConfig().BOTTOKEN
	bot := telegram.NewBot(token)
	bot.AddTextHandler("start", StartCommand{})
	bot.StartPolling()
}
