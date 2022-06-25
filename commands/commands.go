package commands

import (
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

func SetMyBirthdayCommand(b telegram.Bundle) error {
	message := b.Message()

	b.SendMessage(message.GetChatIdStr(), "type your birthday (dd.mm)", true)

	return nil
}

func SetMyBirthdayCommandReply(b telegram.Bundle) error {
	message := b.Message()

	if !isBirthdatyCorrect(b.Message().Text) {
		b.SendMessage(message.GetChatIdStr(), "hmm, i guess there is a typo, try again", true)
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
		b.SendMessage(message.GetChatIdStr(), "Ooops! ", false)
		return err
	}

	b.SendMessage(message.GetChatIdStr(), "Your birthday is: "+user.Birthday, false)

	return nil
}

func isUserInSelectedGroup(b telegram.Bundle) bool {
	return true
}

func saveUserChat(b telegram.Bundle) {}

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

	if isUserInSelectedGroup(b) {
		saveUserChat(b)
	}

	b.SendMessage(
		message.GetChatIdStr(),
		"Hello, i'm celebot! Tell me about your friends birthdays and i will remind you about it ;)",
		true,
	)

	return nil
}

func HelpCommand(b telegram.Bundle) error {
	b.SendMessage(
		b.Message().GetChatIdStr(),
		"I will remind you about birthdays! Here is what i can do..."+"\n\n"+
			"/addme - add your birth day"+"\n"+
			"/addfriend - add your friend's birthday"+"\n"+
			"/mybirthday - show your birthday",
		false,
	)

	return nil
}

func isHenry(username string) bool {
	return true
}

func showBirthdaysFromHenrysClub(b telegram.Bundle) {

}

func ListFromHenrysClub(b telegram.Bundle) error {
	message := b.Message()
	if isHenry(message.From.Username) {
		showBirthdaysFromHenrysClub(b)
	}
	return nil
}
