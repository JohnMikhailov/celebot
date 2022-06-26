package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/meehighlov/celebot/app"
	"github.com/meehighlov/celebot/app/db"
	"github.com/meehighlov/celebot/telegram"
)

func getButtonsRow(texts []string) []telegram.KeyboardButton {
	buttons := make([]telegram.KeyboardButton, 3)
	for _, month := range texts {
		button := telegram.KeyboardButton{Text: month}
		buttons = append(buttons, button)
	}
	return buttons
}

func getKeyboardWithMonths() [][]telegram.KeyboardButton {
	// вынести создание клавиатуры в либу
	return [][]telegram.KeyboardButton{
		getButtonsRow([]string{"jan", "feb", "mar", "apr"}),
		getButtonsRow([]string{"may", "jun", "jul", "aug"}),
		getButtonsRow([]string{"sep", "oct", "nov", "dec"}),
	}
}

func getKeyboardWithDigits() [][]telegram.KeyboardButton {
	// вынести создание клавиатуры в либу
	return [][]telegram.KeyboardButton{
		getButtonsRow([]string{"0", "1", "2", "3"}),
		getButtonsRow([]string{"4", "5", "6", "7"}),
		getButtonsRow([]string{"8", "9", "."}),
	}
}

func SetBirthdayCommand(b telegram.Bundle) error {
	message := b.Message()

	b.SendMessage(message.GetChatIdStr(), "Send me your birthday in format: dd.mm, for example 03.01", true)

	return nil
}

func SetMyBirthdayCommandReply(b telegram.Bundle) error {
	message := b.Message()

	if !isBirthdatyCorrect(b.Message().Text) {
		b.SendMessage(message.GetChatIdStr(), "Hmm, i guess there is a typo, try again please", true)
		return nil
	}

	user := db.User{ID: message.From.Id}
	user.GetById(false)

	user.Birthday = message.Text
	user.Update()

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

	return true
}

func GetBirthDay(b telegram.Bundle) error {
	message := b.Message()
	if !app.IsAllowedUser(message.From.Username) {
		return nil
	}

	user := db.User{ID: message.From.Id}
	err := user.GetById(false)

	if err != nil {
		b.SendMessage(message.GetChatIdStr(), "Hmm, i can't find your birthday... /help", false)
		return err
	}

	b.SendMessage(message.GetChatIdStr(), "Your birthday is: " + user.Birthday, false)

	return nil
}

func StartCommand(b telegram.Bundle) error {
	message := b.Message()
	if !app.IsAllowedUser(message.From.Username) {
		return nil
	}

	user := db.User{
		ID:         message.From.Id,
		Name:       message.From.FirstName,
		TGusername: message.From.Username,
		ChatId:     message.Chat.Id,
	}

	user.Save()

	b.SendMessage(
		message.GetChatIdStr(),
		"Hi, i'm celebot, i will remind you about your frined's birthdays! /help",
		true,
	)

	return nil
}

func showHelpMessage(b telegram.Bundle) {
	b.SendMessage(
		b.Message().GetChatIdStr(),
		"That's what i can do"+"\n\n"+
			"/setme - set your birthday"+"\n"+
			"/addfriend - add your friend's birthday"+"\n"+
			"/mybirthday - show your birthday"+"\n"+
			"/chatbirthdays - show birthdays in chats you own"+"\n"+
			"/hideme - hide your birthday from chat owners"+"\n"+
			"/showme - allow chat owners see yor birthday",
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
		if mem.IsBot && mem.Username == "test_celebot" {
			return true
		}
	}
	return false
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

func ProcessGroupJoin(b telegram.Bundle) error {
	saveGroup(b)

	return nil
}

func DefaultHandler(b telegram.Bundle) error {
	if isBotAddedToGroupEvent(b) {
		return ProcessGroupJoin(b)
	}
	showHelpMessage(b)
	return nil
}

func ShowChatBirthdays(b telegram.Bundle) error {
	message := b.Message()
	chats, err := db.GetUserOwnedChats(message.From.Id)
	if err != nil {
		b.SendMessage(message.GetChatIdStr(), "Ooops, there is a problem occured, i'm working on it...", false)
		return err
	}

	if len(*chats) == 0 {
		b.SendMessage(message.GetChatIdStr(), "I didn't find any chats you own where i was added", false)
		return nil
	}

	chatsBirthdays := ""
	for _, chat := range *chats {
		chatsBirthdays += "For chat " + chat.Title + "\n"

		chatMembers, err := db.GetChatMembers(chat.ID)
		if err != nil {
			log.Panicln("oops")
			continue
		}

		for _, chatMember := range *chatMembers {
			chatsBirthdays += chatMember.Name + " " + chatMember.GetTGUserName() + " " + chatMember.Birthday + "\n"
		}
	}

	b.SendMessage(message.GetChatIdStr(), chatsBirthdays, false)

	return nil
}
