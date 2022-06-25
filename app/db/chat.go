package db

import "log"


func (chat *Chat) Save() error {
	stmt, err := Client.Prepare("INSERT INTO chat(id, type, title, username, firstname, lastname, ownerid) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;")
    if err != nil {
        log.Println("Error when trying to prepare statement for saving chat")
        return err
    }
    defer stmt.Close()

    insertErr := stmt.QueryRow(chat.ID, chat.Type, chat.Title, chat.Username, chat.FirstName, chat.LastName, chat.OwnerId).Scan(&chat.ID)
    if insertErr != nil {
        log.Println("Error when trying to save chat")
        return err
    }
    log.Println("Chat added")

    return nil
}
