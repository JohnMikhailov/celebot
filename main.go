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
	bot.AddEventHandler("/start", &commands.StartCommand{})
	bot.AddEventHandler("/me", &commands.ShowMeCommand{})
	bot.AddEventHandler("/add", &commands.AddFriendCommand{})
	bot.AddEventHandler("/friends", &commands.GetAllFriendsCommand{})
	bot.AddEventHandler("/addme", &commands.AddMyBirthdayCommand{})
	bot.StartPolling()
}
