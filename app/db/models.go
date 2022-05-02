package db

type UserId int64


type User struct {
	// telegram user -> bot's user

	ID int `json:"id"`  // id will be taken from telegram
	Name string `json:"name"`
	TGusername string `json:"tgusername"`

	Friends []Friend
}

type Friend struct {
	ID string `json:"id"`  // uuid
	Name string `json:"name"`
	UserId int `json:"userid"`
	BirthDay string `json:"birthday"`

	Links []Link
}

type Link struct {
	ID int `json:"id"`
	URL string `json:"url"`
	FriendId string `json:"friendid"`
}
