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
	bot.AddHandler("/mybirthday", commands.GetBirthDay)
	bot.AddHandler("/henrysclub", commands.ListFromHenrysClub)
	bot.AddHandler("/help", commands.HelpCommand)

	bot.AddHandler("/setbirthday", commands.SetMyBirthdayCommand)
	bot.AddReplyHandler("type your birthday (dd.mm)", commands.SetMyBirthdayCommandReply)
	bot.AddReplyHandler("hmm, i guess there is a typo, try again", commands.SetMyBirthdayCommandReply)

	bot.SetDefaultHandler(commands.DefaultHandler)

	bot.StartPolling()
}
