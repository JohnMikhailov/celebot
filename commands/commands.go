package commands

import (
	"log"
	"strings"
	"github.com/meehighlov/celebot/telegram"
	"github.com/meehighlov/celebot/app/db"
	"github.com/meehighlov/celebot/app"
)


type StartCommand struct {}
type ShowMeCommand struct {}
type AddFriendCommand struct {}
type GetAllFriendsCommand struct {}
type AddMyBirthdayCommand struct {}
type AddMyBirthdayReplyCommand struct {}


func (handler *AddMyBirthdayReplyCommand) Handle(c *telegram.Context) {
	if !app.IsAllowedUser(c.Message.From.Username) {
		return
	}
	user := db.User{ID: c.Message.From.Id}
	user.GetById(false)
	user.Birthday = c.Message.Text
	c.SendMessage("Your birtday is saved!", c.Message.GetChatIdStr(), false)
}

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

func (handler *AddMyBirthdayCommand) Handle(c *telegram.Context) {
	if !app.IsAllowedUser(c.Message.From.Username) {
		return
	}
	keyboard := telegram.ReplyKeyboardMarkup{Keyboard: getKeyboardWithMonths(), Selective: false, OneTimeKeyboard: true}
	c.SendMessageWithKeyboard(c.Message.GetChatIdStr(), "ok! tell me the month", keyboard)
}

func (handler *StartCommand) Handle(c *telegram.Context) {
	if !app.IsAllowedUser(c.Message.From.Username) {
		return
	}

	user := db.User {
		ID: c.Message.From.Id,
		Name: c.Message.From.FirstName,
		TGusername: c.Message.From.Username,
		ChatId: c.Message.Chat.Id,
	}

	user.Save()

	c.SendMessage(
		c.Message.GetChatIdStr(),
		"Hello, i'm celebot! Tell me about your friends birthdays and i will remind you about it ;)",
		false,
	)
}

func (handler *ShowMeCommand) Handle(c *telegram.Context) {
	if !app.IsAllowedUser(c.Message.From.Username) {
		return
	}
	user := db.User{ID: c.Message.From.Id}
	fetchFriends := false
	user.GetById(fetchFriends)

	c.SendMessage(
		c.Message.GetChatIdStr(),
		"your username is: " + user.TGusername,
		false,
	)
}

func (handler *AddFriendCommand) Handle(c *telegram.Context) {
	if !app.IsAllowedUser(c.Message.From.Username) {
		return
	}
	// params := getCommandParams(c.Message.Text)

	// friend := db.Friend{
	// 	Name: params["name"],
	// 	UserId: c.Message.From.Id,
	// 	BirthDay: params["bd"],
	// 	ChatId: c.Message.Chat.Id,
	// }

	// friend.Save()

	c.SendMessage(
		c.Message.GetChatIdStr(),
		"send your birth day in format: dd.mm",
		true,
	)
}

func (handler *GetAllFriendsCommand) Handle(c *telegram.Context) {
	if !app.IsAllowedUser(c.Message.From.Username) {
		return
	}
	user := db.User{ID: c.Message.From.Id}
	user.GetById(true)

	c.SendMessage(
		c.Message.GetChatIdStr(),
		user.FriendsListAsString(),
		false,
	)
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
