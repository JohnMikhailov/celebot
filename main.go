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
	bot.AddHandler("/me", commands.GetBirthDay)
	bot.AddHandler("/syncgroups", commands.SyncGroupsCommand)
	bot.AddHandler("/chats", commands.ShowChatBirthdays)
	bot.AddHandler("/friends", commands.FriendsListCommand)
	bot.AddHandler("/help", commands.HelpCommand)

	bot.AddHandler("/setme", commands.SetBirthdayCommand)
	bot.AddReplyHandler("Send me your birthday in format: dd.mm, for example 03.01", commands.SetMyBirthdayCommandReply)
	bot.AddReplyHandler("Hmm, i guess there is a typo, try again please", commands.SetMyBirthdayCommandReply)

	bot.AddHandler("/addfriend", commands.AddFriendCommand)
	bot.AddReplyHandler("Ok, send me your friend's name", commands.AddFriendSaveNameCommandReply)
	bot.AddReplyHandler("Ok, now send me your friend's birthday in format: dd.mm, for example 03.01", commands.AddFriendBirthdayCommandReply)
	bot.AddReplyHandler("Ooops, i guess it is in wrong format, try again please", commands.AddFriendBirthdayCommandReply)

	bot.SetDefaultHandler(commands.DefaultHandler)

	bot.StartPolling()
}
