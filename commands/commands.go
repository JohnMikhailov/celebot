package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/meehighlov/celebot/app/db"
	"github.com/meehighlov/celebot/telegram"
)


func SetBirthdayCommand(b telegram.Bundle) error {
	message := b.Message()

	b.SendMessage(message.GetChatIdStr(), "Send me your birthday in format: dd.mm, for example 03.01", true)

	return nil
}

func SetMyBirthdayCommandReply(b telegram.Bundle) error {
	message := b.Message()

	if !isBirthdatyCorrect(message.Text) {
		b.SendMessage(message.GetChatIdStr(), "Hmm, i guess there is a typo, try again please", true)
		return nil
	}

	user := db.User{
		ID:         message.From.Id,
		Name:       message.From.FirstName,
		TGusername: message.From.Username,
		ChatId:     message.Chat.Id,
		Birthday:   message.Text,
	}

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

	user := db.User{
		ID:         message.From.Id,
		Name:       message.From.FirstName,
		TGusername: message.From.Username,
		ChatId:     message.Chat.Id,
		Birthday: "not specified",
	}

	user.Save()

	b.SendMessage(
		message.GetChatIdStr(),
		"Hi, i'm celebot, i will remind you about your friend's birthdays! /help",
		false,
	)

	return nil
}

func showHelpMessage(b telegram.Bundle) {
	b.SendMessage(
		b.Message().GetChatIdStr(),
		"That's what i can do"+"\n\n"+
			"/setme - set your birthday"+"\n"+
			"/addfriend - add your friend's birthday"+"\n"+
			"/friends - show friends list"+"\n"+
			"/clear - clear your friends birthday list"+"\n"+
			"/me - show your birthday"+"\n"+
			"/chats - (beta) show birthdays in chats you own"+"\n"+
			"/syncgroups - (beta) synchronize groups which you and celebot are participated in",
		false,
	)
}

func HelpCommand(b telegram.Bundle) error {
	showHelpMessage(b)
	return nil
}

func isBotAddedToGroupEvent(b telegram.Bundle) bool {
	message := b.Message()
	if message.NewChatMembers == nil {
		return false
	}

	for _, mem := range message.NewChatMembers {
		if mem.IsBot && mem.Username == b.Bot().GetName() {
			return true
		}
	}
	return false
}

func isBotKickedFromGroupEvent(b telegram.Bundle) bool {
	message := b.Message()
	return message.HasLeftChatMember() && message.LeftChatMember.IsBot && message.LeftChatMember.Username == b.Bot().GetName()
}

func saveGroup(b telegram.Bundle) {
	messgae := b.Message()
	chat := messgae.Chat

	owner, err := b.GetChatOwner(messgae.GetChatIdStr())
	if err != nil {
		log.Println("Error getting chat owner: " + err.Error())
		return
	}

	newChat := db.Chat{
		ID:        chat.Id,
		Type:      chat.Type,
		Title:     chat.Title,
		Username:  chat.Username,
		FirstName: chat.FirstName,
		LastName:  chat.LastName,
		OwnerId:   owner.User.Id,
	}

	err = newChat.Save()
	if err != nil {
		log.Println("Chat saving failed: " + err.Error())
	}
}

func deleteGroup(b telegram.Bundle) error {
	message := b.Message()
	chatId := message.Chat.Id

	chat := db.Chat{ID: chatId}
	userChat := db.UserChat{ChatId: chatId}

	chat.Delete()
	userChat.Delete()

	return nil
}

func ProcessGroupJoin(b telegram.Bundle) error {
	saveGroup(b)

	return nil
}

func ProcessGroupKick(b telegram.Bundle) error {
	deleteGroup(b)

	return nil
}

func DefaultHandler(b telegram.Bundle) error {
	if b.IsUpdateFromGroup() {
		if isBotAddedToGroupEvent(b) {
			return ProcessGroupJoin(b)
		}
		if isBotKickedFromGroupEvent(b) {
			return ProcessGroupKick(b)
		}
		return nil
	}

	showHelpMessage(b)
	return nil
}

func ShowChatBirthdays(b telegram.Bundle) error {
	message := b.Message()
	chats, err := db.GetUserOwnedGroups(message.From.Id)
	if err != nil {
		b.SendMessage(message.GetChatIdStr(), "Ooops, there is a problem occured, i'm working on it...", false)
		return err
	}

	if len(*chats) == 0 {
		b.SendMessage(message.GetChatIdStr(), "I didn't find any chats you own where i was added", false)
		return nil
	}

	chatsBirthdays := ""
	isErrorOccured := false
	for _, chat := range *chats {
		chatsBirthdays += "For chat " + chat.Title

		chatMembers, err := db.GetChatMembers(chat.ID)
		if err != nil {
			isErrorOccured = true
			continue
		}

		if len(*chatMembers) == 0 {
			chatsBirthdays += " no birthdays found\n"
			continue
		}

		chatsBirthdays += " found:\n"

		for _, chatMember := range *chatMembers {
			chatsBirthdays += chatMember.Name + " " + chatMember.GetTGUserName() + " " + chatMember.Birthday + "\n"
		}
	}

	if isErrorOccured {
		b.SendMessage(message.GetChatIdStr(), "Ooops, there is a problem occured, i'm working on it...", false)
		return err
	}

	b.SendMessage(message.GetChatIdStr(), chatsBirthdays, false)

	return nil
}

func SyncGroupsCommand(b telegram.Bundle) error {
	message := b.Message()
	userId := message.From.Id

	limit, offset := 10, 0

	chats, err := db.GetAllChats(limit, offset)
	errMessage := "Ooops, there is a problem occured, i'm working on it..."
	if err != nil {
		b.SendMessage(message.GetChatIdStr(), errMessage, false)
		return err
	}

	groupsInCommon := []string{}
	for _, chat := range *chats {
		member, err := b.GetChatMember(chat.ID, userId)
		if err != nil {
			b.SendMessage(message.GetChatIdStr(), errMessage, false)
			return err
		}

		userChat := db.UserChat{UserId: member.User.Id, ChatId: chat.ID}
		err = userChat.Save()
		if err != nil {
			b.SendMessage(message.GetChatIdStr(), errMessage, false)
			return err
		}

		groupsInCommon = append(groupsInCommon, chat.Title)
	}

	if len(groupsInCommon) == 0 {
		b.SendMessage(
			message.GetChatIdStr(),
			"We have no groups incommon! /help",
			false,
		)
		return nil
	}

	foundGroups := strings.Join(groupsInCommon[:], "\n")
	b.SendMessage(
		message.GetChatIdStr(),
		"Cool! We have groups incommon: " + "\n" + foundGroups + "\n",
		false,
	)

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

	b.SendMessage(message.GetChatIdStr(), "Cool! Friend " + friend.Name + " saved", false)

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
