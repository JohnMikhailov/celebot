package db


import "strconv"


type User struct {
	// telegram user -> bot's user

	ID int `json:"id"`  // id will be taken from telegram
	Name string `json:"name"`
	TGusername string `json:"tgusername"`
	ChatId int `json:"chatid"`
	Birthday string `json:"birthday"`

	Friends []Friend
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

func (friend Friend) GetChatIdStr() string {
	return strconv.Itoa(friend.ChatId)
}

type Congratulations struct {
	ID int `json:"id"`
	Text string `json:"text"`
}
