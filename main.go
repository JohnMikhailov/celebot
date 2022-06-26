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
	bot.AddHandler("/syncgroups", commands.SyncGroupsCommand)
	bot.AddHandler("/chatbirthdays", commands.ShowChatBirthdays)
	bot.AddHandler("/help", commands.HelpCommand)

	bot.AddHandler("/setme", commands.SetBirthdayCommand)
	bot.AddReplyHandler("Send me your birthday in format: dd.mm, for example 03.01", commands.SetMyBirthdayCommandReply)
	bot.AddReplyHandler("Hmm, i guess there is a typo, try again please", commands.SetMyBirthdayCommandReply)

	bot.SetDefaultHandler(commands.DefaultHandler)

	bot.StartPolling()
}
