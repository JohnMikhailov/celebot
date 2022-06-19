package db


import (
    "log"
    "fmt"
)


func (user *User) Save() error {
	stmt, err := Client.Prepare("INSERT INTO user(id, name, tgusername, chatid) VALUES($1, $2, $3, $4) RETURNING id;")
    if err != nil {
        log.Println("Error when trying to prepare statement for saving user")
        return err
    }
    defer stmt.Close()

    insertErr := stmt.QueryRow(user.ID, user.Name, user.TGusername, user.ChatId).Scan(&user.ID)
    if insertErr != nil {
        log.Println("Error when trying to save user")
        return err
    }
    log.Println("User added")

    return nil
}

func (user *User) GetById(fetchFriends bool) error {
	stmt, err := Client.Prepare("SELECT id, name, tgusername FROM user WHERE id=$1;")
    if err != nil {
        log.Println("Error when trying to prepare statement for getting user by id")
        return err
    }
    defer stmt.Close()

    result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.Name, &user.TGusername); err != nil {
        log.Println("Error when trying to get User by ID")
        return err
    }

	if fetchFriends {
		stmt, err := Client.Prepare("SELECT id, name, birthday, userid, chatid FROM friend WHERE userid=$1;")
		if err != nil {
			log.Println("Error when trying to prepare statement for fetching friends for user")
			return err
		}

		results, err := stmt.Query(user.ID)
		if err != nil {
			log.Println("Error when fetching friends for user")
			return err
		}

		for results.Next() {
			friend := Friend{}
			err := results.Scan(&friend.ID, &friend.Name, &friend.BirthDay, &friend.UserId, &friend.ChatId)
			if err != nil {
				log.Println("Error when fetching friends for user")
				continue
			}
			user.Friends = append(user.Friends, friend)
		}
	}

    return nil
}

func (user *User) GetFriendsByBirthDate(birthDay string, limit, offset int) error {
    stmt, err := Client.Prepare(
        "SELECT id, name, birthday, userid, chatid FROM friend WHERE birthday like $1 LIMIT $2 OFFSET $3;",
    )
    if err != nil {
        log.Println("Error when trying to prepare statement for fetching friends for user")
        return err
    }

    results, err := stmt.Query(birthDay + ".%", limit, offset)
    if err != nil {
        log.Println("Error when fetching friends for user by birthday")
        return err
    }

    for results.Next() {
        friend := Friend{}
        err := results.Scan(&friend.ID, &friend.Name, &friend.BirthDay, &friend.UserId, &friend.ChatId)
        if err != nil {
            log.Println("Error when fetching friends for user by birthday")
            continue
        }
        user.Friends = append(user.Friends, friend)
    }

    return nil
}

func (friend *Friend) Save() error {
	stmt, err := Client.Prepare("INSERT INTO friend(name, birthday, userid, chatid) VALUES($1, $2, $3, $4) RETURNING id;")
    if err != nil {
        log.Println("Error when trying to prepare statement")
        return err
    }
    defer stmt.Close()

    insertErr := stmt.QueryRow(friend.Name, friend.BirthDay, friend.UserId, friend.ChatId).Scan(&friend.ID)
    if insertErr != nil {
        _ = fmt.Sprintf("Error when trying to save friend with id: %d", friend.ID)
        return err
    }
    log.Println("Friend added")

    return nil
}

func (link *Link) Save() error {
	stmt, err := Client.Prepare("INSERT INTO link(url, friendid) VALUES($1, $2);")
	if err != nil {
        log.Println("Error when trying to prepare statement")
        return err
    }
    defer stmt.Close()
	insertErr := stmt.QueryRow(link.URL, link.FriendId)
    if insertErr != nil {
        log.Println("Error when trying to save link")
        return err
    }
    log.Println("Link added")
    return nil
}
