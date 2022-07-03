package db

import "log"


func (chat *Chat) Save() error {
	stmt, err := Client.Prepare(
		"INSERT INTO chat(id, type, title, username, firstname, lastname, ownerid) " +
		"VALUES($1, $2, $3, $4, $5, $6, $7) " +
		"ON CONFLICT(id) DO UPDATE SET type=$2, title=$3, username=$4, firstname=$5, lastname=$6, ownerid=$7 " +
		"RETURNING id;",
	)
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

	log.Println("Chat created/updated")

	return nil
}

func GetUserOwnedGroups(userId int) (*[]Chat, error) {
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
		"SELECT user.id, user.name, user.tgusername, user.chatid, user.birthday FROM user " +
		"INNER JOIN userchat ON user.id = userchat.userid " +
		"INNER JOIN chat ON chat.id = userchat.chatid " +
		"WHERE userchat.chatid = $1;",
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
		err := results.Scan(&user.ID, &user.Name, &user.TGusername, &user.ChatId, &user.Birthday)
		if err != nil {
			log.Println("Error when fetching chat for user: " + err.Error())
			continue
		}
		users = append(users, user)
	}

	return &users, nil
}

func (userChat *UserChat) Save() error {
	stmt, err := Client.Prepare(
		"INSERT INTO userchat(userid, chatid) " +
		"VALUES($1, $2) " +
		"ON CONFLICT(userid, chatid) DO NOTHING;",
	)
	if err != nil {
		log.Println("Error when trying to prepare statement for saving userchat: " + err.Error())
		return err
	}
	defer stmt.Close()

	insertErr := stmt.QueryRow(&userChat.UserId, &userChat.ChatId)
	if insertErr.Err() != nil {
		log.Println("Error when trying to save userchat: " + insertErr.Err().Error())
		return insertErr.Err()
	}
	log.Println("UserChat created/updated")

	return nil
}

func GetAllChats(limit, offset int) (*[]Chat, error) {
	stmt, err := Client.Prepare(
		`SELECT id, type, title, username, firstname, lastname, ownerid FROM chat LIMIT $1 OFFSET $2;`,
	)
	if err != nil {
		log.Println("Error when trying to prepare statement for getting all celebot chats: " + err.Error())
		return nil, err
	}
	defer stmt.Close()

	results, err := stmt.Query(limit, offset)
	if err != nil {
		log.Println("Error when trying to get all celebot chats: " + err.Error())
		return nil, err
	}

	chats := []Chat{}
	for results.Next() {
		chat := Chat{}
		err := results.Scan(&chat.ID, &chat.Type, &chat.Title, &chat.Username, &chat.FirstName, &chat.LastName, &chat.OwnerId)
		if err != nil {
			log.Println("Error when fetching celebot chat: " + err.Error())
			continue
		}
		chats = append(chats, chat)
	}

	return &chats, nil
}

func (chat *Chat) Delete() error {
	stmt, err := Client.Prepare(
		`DELETE FROM chat WHERE id = $1`,
	)
	if err != nil {
		log.Println("Error when trying to prepare statement for deleting chat: " + err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(&chat.ID)
	if err != nil {
		log.Println("Error when trying to delete chat: " + err.Error())
		return err
	}

	log.Println("Chat deleted")

	return nil
}

func (userChat *UserChat) Delete() error {
	stmt, err := Client.Prepare(
		`DELETE FROM userchat WHERE chatid = $1;`,
	)
	if err != nil {
		log.Println("Error when trying to prepare statement for deleting userchat: " + err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(&userChat.ID)
	if err != nil {
		log.Println("Error when trying to delete userchat: " + err.Error())
		return err
	}

	log.Println("UserChat deleted")

	return nil
}
