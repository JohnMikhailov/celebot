package commands

import (
	"strconv"
	"strings"

	"github.com/meehighlov/celebot/app/db"
	"github.com/meehighlov/celebot/app"
	"github.com/meehighlov/celebot/telegram"
)


func SetBirthdayCommand(b telegram.Bundle) error {
	message := b.Message()
	if !app.IsAuthUser(message.From.Id) {
		return nil
	}

	b.SendMessage(message.GetChatIdStr(), "Send me your birthday in format: dd.mm, for example 03.01", true)

	return nil
}

func SetMyBirthdayCommandReply(b telegram.Bundle) error {
	message := b.Message()

	if !isBirthdatyCorrect(message.Text) {
		b.SendMessage(message.GetChatIdStr(), "Hmm, i guess there is a typo, try again please", true)
		return nil
	}

	user := db.User{ID: message.From.Id}
	err := user.Get()
	if err != nil {
		b.SendMessage(message.GetChatIdStr(), "Ooops, there is a problem occured, i'm working on it 😅", false)
		return nil
	}

	user.Birthday = message.Text

	user.Save()

	b.SendMessage(message.GetChatIdStr(), "Cool, i've saved your birthday!", false)

	return nil
}

func isBirthdatyCorrect(birtday string) bool {
	parts := strings.Split(birtday, ".")
	if len(parts) != 2 || (len(parts[0])+len(parts[1])) != 4 {
		return false
	}

	for _, nums := range parts {
		_, err := strconv.Atoi(nums)
		if err != nil {
			return false
		}
	}

	day, _ := strconv.Atoi(parts[0])
	if !(day >= 1 && day <= 31) {
		return false
	}

	month, _ := strconv.Atoi(parts[1])
	if !(month >= 1 && month <= 12) {
		return false
	}

	if day >= 30 && month == 2 {
		// TODO add same checks for other months
		return false
	}

	return true
}

func GetBirthDay(b telegram.Bundle) error {
	message := b.Message()
	if !app.IsAuthUser(message.From.Id) {
		return nil
	}

	user := db.User{ID: message.From.Id}
	err := user.Get()

	if err != nil {
		b.SendMessage(message.GetChatIdStr(), "Hmm, i can't find your birthday... /help", false)
		return err
	}

	b.SendMessage(message.GetChatIdStr(), "Your birthday is " + user.Birthday, false)

	return nil
}

func StartCommand(b telegram.Bundle) error {
	message := b.Message()

	b.SendMessage(
		message.GetChatIdStr(),
		"Hi, i'm celebot, i will remind you about your friend's birthdays!" + "\n" +
		"Got access code? Call /code to pass it!",
		false,
	)

	return nil
}

func showHelpMessage(b telegram.Bundle, additional *string) {
	helpMessage := (
		"That's what i can do"+"\n\n"+
		"/me - show your birthday"+"\n"+
		"/setme - set your birthday"+"\n"+
		"/addfriend - add your friend's birthday"+"\n"+
		"/friends - show friends list"+"\n"+
		"/clear - clear your friends birthday list"+"\n"+
		"/code - use this command to pass access code")

	if additional != nil {
		helpMessage += "\n" + *additional
	}

	b.SendMessage(
		b.Message().GetChatIdStr(),
		helpMessage,
		false,
	)
}

func showHelpMessageNoAuth(b telegram.Bundle) {
	b.SendMessage(
		b.Message().GetChatIdStr(),
		"/code - use this command to pass access code",
		false,
	)
}

func showHelpMessageAdmin(b telegram.Bundle) {
	message := "/chat - show all birthdays in chat"
	showHelpMessage(b, &message)
}

func HelpCommand(b telegram.Bundle) error {
	message := b.Message()

	if !app.IsAuthUser(message.From.Id) {
		showHelpMessageNoAuth(b)
		return nil
	}

	if app.IsAdmin(message.From.Id) {
		showHelpMessageAdmin(b)
		return nil
	}

	showHelpMessage(b, nil)
	return nil
}

func DefaultHandler(b telegram.Bundle) error {
	if b.IsUpdateFromGroup() {
		return nil
	}

	return nil
}

func AddFriendCommand(b telegram.Bundle) error {
	chatId := b.Message().GetChatIdStr()
	b.SendMessage(chatId, "Ok, send me your friend's name", true)
	return nil
}

func AddFriendSaveNameCommandReply(b telegram.Bundle) error {
	message := b.Message()
	friend := db.Friend{
		Name: message.Text,
		BirthDay: "not specified",
		UserId: message.From.Id,
		ChatId: message.Chat.Id,
	}
	friend.Save()

	b.SendMessage(message.GetChatIdStr(), "Ok, now send me your friend's birthday in format: dd.mm, for example 03.01", true)

	return nil
}

func AddFriendBirthdayCommandReply(b telegram.Bundle) error {
	message := b.Message()

	if !isBirthdatyCorrect(message.Text) {
		b.SendMessage(message.GetChatIdStr(), "Ooops, i guess it is in wrong format, try again please", true)
		return nil
	}

	friend := db.Friend{
		UserId: message.From.Id,
		ChatId: message.Chat.Id,
	}

	friend.GetFriendWithUnspecifiedBirthday()

	friend.UpdateForBirthday(friend.Name, message.Text)

	b.SendMessage(message.GetChatIdStr(), "Cool! Friend " + friend.Name + " saved 😉", false)

	return nil
}

func FriendsListCommand(b telegram.Bundle) error {
	message := b.Message()

	user := db.User{ID: message.From.Id}

	user.GetWithFriends(true)
	friend := db.Friend{UserId: user.ID}

	friend.DeleteEmptyBirthdays()

	friendsList := user.FriendsListAsString()
	if friendsList == "" {
		friendsList = "Friends list is empty! /help"
	}

	b.SendMessage(message.GetChatIdStr(), friendsList, false)

	return nil
}

func ClearFriendsListCommand(b telegram.Bundle) error {
	message := b.Message()
	if !app.IsAuthUser(message.From.Id) {
		return nil
	}

	b.SendMessage(message.GetChatIdStr(), "A you sure you want to clear friends list? Send any key", true)
	return nil
}

func ClearFriendsListReplyCommand(b telegram.Bundle) error {
	message := b.Message()

	friend := db.Friend{UserId: message.From.Id}
	friend.DeleteFriendsByUserId()

	b.SendMessage(message.GetChatIdStr(), "Friends list is clear!", false)

	return nil
}

func saveClubUser(b telegram.Bundle, isAdmin bool) error {
	message := b.Message()

	dbIsAdmin := 0
	if isAdmin {
		dbIsAdmin = 1
	}

	user := db.User{
		ID:         message.From.Id,
		Name:       message.From.FirstName,
		TGusername: message.From.Username,
		ChatId:     message.Chat.Id,
		Birthday:   "not specified",
		IsAdmin:    dbIsAdmin,
	}

	user.Save()

	return nil
}

func AuthCodeCommand(b telegram.Bundle) error {
	b.SendMessage(b.Message().GetChatIdStr(), "Enter access code", true)
	return nil
}

func AuthCodeCommandReply(b telegram.Bundle) error {
	message := b.Message()

	code := message.Text
	clubCode := app.GetConfig().CLUBCODE
	adminCode := app.GetConfig().ADMINCODE

	switch code {
	case clubCode:
		isAdmin := false
		saveClubUser(b, isAdmin)
	case adminCode:
		isAdmin := true
		saveClubUser(b, isAdmin)
	default:
		b.SendMessage(message.GetChatIdStr(), "Incorrect code 🙂", false)
		return nil
	}

	b.SendMessage(message.GetChatIdStr(), "You rock 😎 Now command list is available for you /help", false)

	return nil
}

func ChatCommand(b telegram.Bundle) error {
	message := b.Message()
	if !app.IsAdmin(message.From.Id) {
		return nil
	}

	users, err := db.GetAllClubUsers()
	if err != nil {
		b.SendMessage(message.GetChatIdStr(), "Ooops, there is a problem occured, i'm working on it 😅", false)
		return err
	}

	text := ""
	for _, user := range *users {
		text += user.Name + " " + user.GetTGUserName() + " " + user.Birthday
	}

	b.SendMessage(message.GetChatIdStr(), text, false)

	return nil
}
