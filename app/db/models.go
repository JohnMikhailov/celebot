package db


type User struct {
	// telegram user -> bot's user

	ID int `json:"id"`  // id will be taken from telegram
	Name string `json:"name"`
	TGusername string `json:"tgusername"`
	ChatId int `json:"chatid"`

	Friends []Friend
}

func (user User) FriendsListAsString() string {
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

	Links []Link
}

type Link struct {
	ID int `json:"id"`
	URL string `json:"url"`
	FriendId string `json:"friendid"`
}

type Congratulations struct {
	ID int `json:"id"`
	Text string `json:"text"`
}
