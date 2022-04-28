package telegram


type Chat struct {
	// full description https://core.telegram.org/bots/api#chat

	Id int `json:"id"`

	//Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string  `json:"type"`

	Title string  `json:"title"`
	Username string  `json:"username"`

}

type User struct {
	// full description https://core.telegram.org/bots/api#user

	Id int  `json:"id"`
	IsBot bool  `json:"is_bot"`
	FirstName string  `json:"first_name"`
	LastName string  `json:"last_name"`
	Username string  `json:"username"`

}

type Message struct {
	Message_id int  `json:"message_id"`
	From User  `json:"from"`
	SenderChat Chat  `json:"sender_chat"`
	Chat Chat `json:"chat"`
	Text string  `json:"text"`
}


type Update struct {
	UpdateId int `json:"update_id"`
	Message Message `json:"message"`
}


type UpdateResponse struct {
	Ok bool `json:"ok"`
	Result []Update `json:"result"`
}


type ResponseParameters struct {
	MigrateToChatId int `json:"migrate_to_chat_id"`
	RetryAfter int `json:"retry_after"`
}


type ErrorResponse struct {
	Ok bool `json:"ok"`
	Description string `json:"description"`
	Parameters ResponseParameters  `json:"parameters"`
}
