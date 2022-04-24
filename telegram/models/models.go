package telegram


type Chat struct {
	// full description https://core.telegram.org/bots/api#chat

	id int `json:"id"`

	//Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	type_ string  `json:"type"`

	title string  `json:"title"`
	username string  `json:"username"`

}

type User struct {
	// full description https://core.telegram.org/bots/api#user

	id int  `json:"id"`
	is_bot bool  `json:"is_bot"`
	first_name string  `json:"first_name"`
	last_name string  `json:"last_name"`
	username string  `json:"username"`

}

type Message struct {
	message_id int  `json:"message_id"`
	from User  `json:"from"`
	sender_chat Chat  `json:"sender_chat"`
	text string  `json:"text"`
}
