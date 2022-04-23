package telegram


type Chat struct {
	// full description https://core.telegram.org/bots/api#chat

	id int

	//Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	type_ string

	title string
	username string

}
