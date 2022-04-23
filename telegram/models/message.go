package telegram

import (
	"telegram/models/user"
	"telegram/models/chat"
)



type Message struct {
	message_id int
	from user.User
	sender_chat chat.Chat
}
