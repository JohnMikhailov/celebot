package db


import "strconv"


type User struct {
	// telegram user -> bot's user

	ID int `json:"id"`  // id will be taken from telegram
	Name string `json:"name"`
	TGusername string `json:"tgusername"`
	ChatId int `json:"chatid"`  // chatId - id of chat with user, bot uses it to send notification
	Birthday string `json:"birthday"`
	ShowToChatOwner int `json:"showtochatowner"`

	Friends []Friend
}

func (user *User) GetTGUserName() string {
	if string(user.TGusername[0]) == "@" {
		return user.TGusername
	}
	return "@" + user.TGusername
}

func (user *User) FriendsListAsString() string {
	result := ""
	for _, friend := range user.Friends {
		result += friend.Name + " " + friend.BirthDay + "\n"
	}
	return result
}

type Friend struct {
	ID int `json:"id"`  // uuid
	Name string `json:"name"`
	UserId int `json:"userid"`
	BirthDay string `json:"birthday"`
	ChatId int `json:"chatid"`
}

type Chat struct {
	ID int `json:"id"`
	Type string `json:"type"`
	Title string `json:"title"`
	Username string `json:"username"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	OwnerId int `json:"ownerid"`
}

type UserChat struct {
	UserId int `json:"userid"`
	ChatId int `json:"chatid"`
}

func (friend Friend) GetChatIdStr() string {
	return strconv.Itoa(friend.ChatId)
}
