package db

type userId int64

type User struct {
	ID userId `json:"id"`
	Name string `json:"name"`
	BirthDate string `json:"birthDate"`  // sqlite3 not support Date type

	UserLinks []UserLink
}

type UserLink struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`  // example: https://social.media.domain
	UserId userId `json:"userId"`
}
