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

func getKeyboardWithMonths() [][]telegram.KeyboardButton{
	// вынести создание клавиатуры в либу
	return [][]telegram.KeyboardButton{
		getButtonsRow([]string{"jan", "feb", "mar", "apr"}),
		getButtonsRow([]string{"may", "jun", "jul", "aug",}),
		getButtonsRow([]string{"sep", "oct", "nov", "dec"}),
	}
}

func AddMyBirthdayCommand(b telegram.Bundle) error {
	message := b.Message()
	if !app.IsAllowedUser(message.From.Username) {
		return nil
	}
	keyboard := telegram.ReplyKeyboardMarkup{Keyboard: getKeyboardWithMonths(), Selective: false, OneTimeKeyboard: true}
	b.SendMessageWithKeyboard(message.GetChatIdStr(), keyboard)

	return nil
}

func isBirthdatyCorrect(birtday string) bool {
	parts := strings.Split(birtday, ".")
	if len(parts) != 2 || (len(parts[0]) + len(parts[1])) != 4 {
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

func SaveBirthdayWithArgs(b telegram.Bundle) error {
	message := b.Message()

	birthDayArgs := b.Args()
	if len(birthDayArgs) != 1 || !isBirthdatyCorrect(birthDayArgs[0]) {
		b.SendMessage(
			message.GetChatIdStr(),
			"Hmm, i guess there is some typo... try again" + "\n" +
			"Stuck? Call /help for commands description",
			false,
		)
		return nil
	}

	birthDay := birthDayArgs[0]

	user := db.User {ID: message.From.Id}
	user.GetById(false)

	user.Birthday = birthDay
	user.Update()

	b.SendMessage(
		message.GetChatIdStr(),
		"I saved your birthday! It is " + message.Text + "\n" +
		" - if you made a mistake, please call /addme again",
		false,
	)

	return nil
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

	b.SendMessage(message.GetChatIdStr(), "Your birthday is: " + user.Birthday, false)

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

	user := db.User {
		ID: message.From.Id,
		Name: message.From.FirstName,
		TGusername: message.From.Username,
		ChatId: message.Chat.Id,
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

func ShowMeCommand(b telegram.Bundle) error {
	message := b.Message()
	if !app.IsAllowedUser(message.From.Username) {
		return nil
	}
	user := db.User{ID: message.From.Id}
	fetchFriends := false
	user.GetById(fetchFriends)

	b.SendMessage(
		message.GetChatIdStr(),
		"your username is: " + user.TGusername,
		true,
	)

	return nil
}

func HelpCommand(b telegram.Bundle) error {
	b.SendMessage(
		b.Message().GetChatIdStr(),
		"I will remind you about birthdays! Here is what i can do..." + "\n\n" +
		"/addme - add your birth day" + "\n" +
		"/addfriend - add your friend's birthday" + "\n" +
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
