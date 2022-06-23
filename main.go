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
	bot.AddHandler("/addme", commands.SaveBirthdayWithArgs)
	bot.AddHandler("/mybirthday", commands.GetBirthDay)
	bot.AddHandler("/henrysclub", commands.ListFromHenrysClub)
	bot.AddHandler("/help", commands.HelpCommand)

	bot.AddHandler("/add", commands.AddMyBirthdayCommand)
	bot.AddReplyHandler("type your birthday (dd.mm)", commands.AddMyBirthdayCommandReply)
	bot.AddReplyHandler("hmm, i guess there is a typo, try again", commands.AddMyBirthdayCommandReply)

	bot.StartPolling()
}
