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
	token := app.GetConfig().BOTTOKEN
	bot := telegram.NewBot(token)
	bot.AddEventHandler("/start", commands.StartCommand{})
	bot.AddEventHandler("congrats", commands.RandomCongratulationCommand{})
	bot.AddEventHandler("me", commands.ShowMeCommand{})
	bot.StartPolling()
}
