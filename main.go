package main

import (
	//"os"
	//"io/ioutil"
	//"net/http"
	"github.com/meehighlov/celebot/telegram"
	"github.com/meehighlov/celebot/commands"
	"github.com/meehighlov/celebot/app"
)

func main() {
	token := app.getConfig().BOTTOKEN
	bot := telegram.NewBot(token)
	bot.AddEventHandler("/start", commands.StartCommand{})
	bot.AddEventHandler("/add", commands.AddPersonCommand{})
	bot.StartPolling()
}
