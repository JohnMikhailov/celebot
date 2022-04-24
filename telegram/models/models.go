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
	Is_bot bool  `json:"is_bot"`
	First_name string  `json:"first_name"`
	Last_name string  `json:"last_name"`
	Username string  `json:"username"`

}

type Message struct {
	Message_id int  `json:"message_id"`
	From User  `json:"from"`
	Sender_chat Chat  `json:"sender_chat"`
	Text string  `json:"text"`
}


type Update struct {
	Update_id int `json:"update_id"`
	MessageInfo Message `json:"message"`
}


type UpdateResponse struct {
	Ok bool `json:"ok"`
	Result []Update `json:"result"`
}
