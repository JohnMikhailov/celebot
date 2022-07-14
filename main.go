package main

import (
	"github.com/meehighlov/celebot/app"
	"github.com/meehighlov/celebot/commands"
	"github.com/meehighlov/celebot/telegram"
)

func main() {
	logFileName := app.GetConfig().LOG_FILE
	logFile := app.SetupFileLogging(logFileName)
	defer logFile.Close()

	token := app.GetConfig().BOTTOKEN_CELEBOT
	commands.RunChecks()
	bot := telegram.NewBot(token)
	bot.AddHandler("/start", commands.StartCommand)
	bot.AddHandler("/me", commands.GetBirthDay)
	bot.AddHandler("/help", commands.HelpCommand)
	bot.AddHandler("/show", commands.FriendsListCommand)
	bot.AddHandler("/chat", commands.ChatCommand)

	bot.AddHandler("/setme", commands.SetBirthdayCommand)
	bot.AddReplyHandler("Send me your birthday in format: dd.mm, for example 03.01", commands.SetMyBirthdayCommandReply)
	bot.AddReplyHandler("Hmm, i guess there is a typo, try again please", commands.SetMyBirthdayCommandReply)

	bot.AddHandler("/add", commands.AddFriendCommand)
	bot.AddReplyHandler("Ok, send me your friend's name", commands.AddFriendSaveNameCommandReply)
	bot.AddReplyHandler("Ok, now send me your friend's birthday in format: dd.mm, for example 03.01", commands.AddFriendBirthdayCommandReply)
	bot.AddReplyHandler("Ooops, i guess it is in wrong format, try again please", commands.AddFriendBirthdayCommandReply)

	bot.AddHandler("/clear", commands.ClearFriendsListCommand)
	bot.AddReplyHandler("A you sure you want to clear friends list? Send any key", commands.ClearFriendsListReplyCommand)

	bot.AddHandler("/code", commands.AuthCodeCommand)
	bot.AddReplyHandler("Enter access code", commands.AuthCodeCommandReply)

	bot.SetDefaultHandler(commands.DefaultHandler)

	bot.StartPolling()
}
