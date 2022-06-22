package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/meehighlov/celebot/app"
	"github.com/meehighlov/celebot/app/db"
	"github.com/meehighlov/celebot/telegram"
)

// func (handler *AddMyBirthdayReplyCommand) Handle(c *telegram.Context) {
// 	if !app.IsAllowedUser(c.Message.From.Username) {
// 		return
// 	}
// 	user := db.User{ID: c.Message.From.Id}
// 	user.GetById(false)
// 	user.Birthday = c.Message.Text
// 	c.SendMessage("Your birtday is saved!", c.Message.GetChatIdStr(), false)
// }

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
	b.SendMessageWithKeyboard(message.GetChatIdStr(), "ok! tell me the month", keyboard)

	return nil
}

func SaveBirthdayCommand(b telegram.Bundle) error {
	message := b.Message()
	firstStepText := "Send your birthday in format: dd.mm, for example 01.03 (march third)"
	errorStepText := "Ooops! I think birthday is not in format: dd.mm"
	if !app.IsAllowedUser(message.From.Username) {
		return nil
	}

	log.Println("reply text: " + message.ReplyToMessage.Text)

	switch message.ReplyToMessage.Text {
		case firstStepText:
			SaveBirthday(b)
			return nil
		case errorStepText:
			b.SendMessage(message.GetChatIdStr(), firstStepText)
			return nil
	}

	b.SendMessage(message.GetChatIdStr(), firstStepText)

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

func SaveBirthday(b telegram.Bundle) error {
	message := b.Message()
	errorStepText := "Ooops! I think birthday is not in format: dd.mm"
	if !isBirthdatyCorrect(message.Text) {
		b.SendMessage(message.GetChatIdStr(), errorStepText)
		return nil
	}

	user := db.User {ID: message.From.Id}
	user.GetById(false)

	user.Birthday = message.Text
	user.Save()

	b.SendMessage(
		message.GetChatIdStr(),
		"I saved your birthday! It is " + "message.Text" + "\n" +
		" - if you made a mistake, please call /addme again",
	)

	return nil
}

func GetBirthDay(b telegram.Bundle) error {
	message := b.Message()
	if !app.IsAllowedUser(message.From.Username) {
		return nil
	}

	user := db.User {ID: message.From.Id}
	err := user.GetById(false)

	if err != nil {
		b.SendMessage(message.GetChatIdStr(), "Hmm, it seems like i don't know your birhday yet... try to call /addme")
		return err
	}

	b.SendMessage(message.GetChatIdStr(), "Your birthday is: " + user.Birthday)

	return nil
}

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

	b.SendMessage(
		message.GetChatIdStr(),
		"Hello, i'm celebot! Tell me about your friends birthdays and i will remind you about it ;)",
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
	)

	return nil
}

func getCommandParams(text string) map[string]string {
	// command syntax: command param1=value1 param2=value2
	log.Println("raw message text:", text)
	trancatedCommand := strings.Fields(text)

	var preparedParams = map[string]string{}

	params := trancatedCommand[1:]

	for _, param := range params {
		splitedParam := strings.Split(param, "=")
		if len(splitedParam) > 1 {
			paramName := splitedParam[0]
			paramValue := splitedParam[1]
			preparedParams[paramName] = paramValue
		}
	}

	log.Println("prepared params:", preparedParams)

	return preparedParams
}
