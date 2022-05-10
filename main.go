package main

import (
	"github.com/meehighlov/celebot/telegram"
	"github.com/meehighlov/celebot/commands"
	"github.com/meehighlov/celebot/app"
)

func main() {
	token := app.GetConfig().BOTTOKEN
	go commands.CheckBirthDays()
	bot := telegram.NewBot(token)
	bot.AddEventHandler("/start", &commands.StartCommand{})
	bot.AddEventHandler("congrats", &commands.RandomCongratulationCommand{})
	bot.AddEventHandler("me", &commands.ShowMeCommand{})
	bot.AddEventHandler("add", &commands.AddFriendCommand{})
	bot.AddEventHandler("friends", &commands.GetAllFriendsCommand{})
	bot.StartPolling()
}
