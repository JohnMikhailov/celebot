package db

import "log"


func (chat *Chat) Save() error {
	stmt, err := Client.Prepare("INSERT INTO chat(id, type, title, username, firstname, lastname, ownerid) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;")
    if err != nil {
        log.Println("Error when trying to prepare statement for saving chat: " + err.Error())
        return err
    }
    defer stmt.Close()

    insertErr := stmt.QueryRow(chat.ID, chat.Type, chat.Title, chat.Username, chat.FirstName, chat.LastName, chat.OwnerId).Scan(&chat.ID)
    if insertErr != nil {
        log.Println("Error when trying to save chat: " + err.Error())
        return err
    }
    log.Println("Chat added")

    return nil
}

func GetUserOwnedChats(userId int) (*[]Chat, error) {
	stmt, err := Client.Prepare("SELECT id, type, title, username, firstname, lastname, ownerid FROM chat WHERE ownerid = $1;")
    if err != nil {
        log.Println("Error when trying to prepare statement for getting chats: " + err.Error())
        return nil, err
    }
    defer stmt.Close()

    results, err := stmt.Query(userId)
    if err != nil {
        log.Println("Error when trying to get owner's chats: " + err.Error())
        return nil, err
    }

	chats := []Chat{}
	for results.Next() {
		chat := Chat{}
		err := results.Scan(&chat.ID, &chat.Type, &chat.Title, &chat.Username, &chat.FirstName, &chat.LastName, &chat.OwnerId)
		if err != nil {
			log.Println("Error when fetching chat for user")
			continue
		}
		chats = append(chats, chat)
	}

    return &chats, nil
}

func GetChatMembers(chatId int) (*[]User, error) {
	stmt, err := Client.Prepare(
		`SELECT id, name, tgusername, chatid, birthday, showtochatowner FROM
		   user
		   INNER JOIN userchat ON user.id = userchat.userid
		   INNER JOIN chat ON chat.id = userchat.chatid
		   WHERE chatid = $1 AND user.showtochatowner = 1;
		`,
	)
    if err != nil {
        log.Println("Error when trying to prepare statement for getting chat members: " + err.Error())
        return nil, err
    }
    defer stmt.Close()

	results, err := stmt.Query(chatId)
    if err != nil {
        log.Println("Error when trying to get chat members: " + err.Error())
        return nil, err
    }

	users := []User{}
	for results.Next() {
		user := User{}
		err := results.Scan(&user.ID, &user.Name, &user.TGusername, &user.ChatId, &user.Birthday, user.ShowToChatOwner)
		if err != nil {
			log.Println("Error when fetching chat for user")
			continue
		}
		users = append(users, user)
	}

	return &users, nil
}
