package commands

import (
	"log"
	"fmt"
	"strings"
	"github.com/meehighlov/celebot/telegram"
	"github.com/meehighlov/celebot/app/db"
)


type StartCommand struct {}
type AddPersonCommand struct {}
type RandomCongratulationCommand struct {}
type ShowMeCommand struct {}
type AddFriendCommand struct {}
type GetAllFriendsCommand struct {}


func (handler *StartCommand) Handle(c *telegram.Context) {
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
	)
}

func (handler *AddPersonCommand) Handle(c *telegram.Context) {
	params := getCommandParams(c.Message.Text)
	name := params["name"]
	bd := params["bd"]

	text := fmt.Sprintf("Added new person: %s birth date: %s", name, bd)
	c.SendMessage(
		c.Message.GetChatIdStr(),
		text,
	)
}

func (handler *RandomCongratulationCommand) Handle(c *telegram.Context) {
	c.SendMessage(
		c.Message.GetChatIdStr(),
		"i don't know any congratulations yet, may be you would like add one?:)",
	)
}

func (handler *ShowMeCommand) Handle(c *telegram.Context) {
	user := db.User{ID: c.Message.From.Id}
	fetchFriends := false
	user.GetById(fetchFriends)

	c.SendMessage(
		c.Message.GetChatIdStr(),
		"your username is: " + user.TGusername,
	)
}

func (handler *AddFriendCommand) Handle(c *telegram.Context) {
	params := getCommandParams(c.Message.Text)

	friend := db.Friend{
		Name: params["name"],
		UserId: c.Message.From.Id,
		BirthDay: params["bd"],
		ChatId: c.Message.Chat.Id,
	}

	friend.Save()

	c.SendMessage(
		c.Message.GetChatIdStr(),
		"new Birth Day added for friend: " + friend.Name,
	)
}

func (handler *GetAllFriendsCommand) Handle(c *telegram.Context) {
	user := db.User{ID: c.Message.From.Id}
	user.GetById(true)

	c.SendMessage(
		c.Message.GetChatIdStr(),
		user.FriendsListAsString(),
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
