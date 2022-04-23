package telegram


type User struct {
	// full description https://core.telegram.org/bots/api#user

	id int
	is_bot bool
	first_name string
	last_name string
	username string

}