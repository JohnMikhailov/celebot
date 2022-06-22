package main

import (
	"github.com/meehighlov/celebot/app"
	"github.com/meehighlov/celebot/commands"
	"github.com/meehighlov/celebot/telegram"
)

func main() {
	token := app.GetConfig().BOTTOKEN_CELEBOT
	commands.RunChecks()
	bot := telegram.NewBot(token)
	bot.AddHandler("/start", commands.StartCommand)
	bot.AddHandler("/addme", commands.SaveBirthdayCommand)
	bot.AddHandler("/mybirthday", commands.GetBirthDay)
	bot.StartPolling()
}
