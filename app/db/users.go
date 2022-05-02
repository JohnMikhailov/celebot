package db


import "fmt"


func (user User) Save() error {
	stmt, err := Client.Prepare("INSERT INTO user(id, name, tgusername) VALUES($1, $2, $3);")
    if err != nil {
        fmt.Println("Error when trying to prepare statement")
        fmt.Println(err)
        return err
    }
    defer stmt.Close()

    insertErr := stmt.QueryRow(user.ID, user.Name, user.TGusername)
    if insertErr != nil {
        fmt.Println("Error when trying to save user")
		fmt.Println(insertErr)
        return err
    }
    fmt.Println("User added")

    return nil
}

func (user User) GetById() error {
	stmt, err := Client.Prepare("SELECT id, name, tgusername FROM user WHERE id=$1;")
    if err != nil {
        fmt.Println("Error when trying to prepare statement")
        fmt.Println(err)
        return err
    }
    defer stmt.Close()

    result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.Name, &user.TGusername); err != nil {
        fmt.Println("Error when trying to get User by ID")
        return err
    }

    return nil
}

func (friend Friend) Save() error {
	stmt, err := Client.Prepare("INSERT INTO user(id, name, brithday, userid) VALUES($1, $2, $3);")
    if err != nil {
        fmt.Println("Error when trying to prepare statement")
        fmt.Println(err)
        return err
    }
    defer stmt.Close()

    insertErr := stmt.QueryRow(friend.ID, friend.Name, friend.BirthDay, friend.UserId)
    if insertErr != nil {
        fmt.Println("Error when trying to save user")
        return err
    }
    fmt.Println("Friend added")

    return nil
}

func (link Link) Save() error {
	stmt, err := Client.Prepare("INSERT INTO link(url, friendid) VALUES($1, $2);")
	if err != nil {
        fmt.Println("Error when trying to prepare statement")
        fmt.Println(err)
        return err
    }
    defer stmt.Close()
	insertErr := stmt.QueryRow(link.URL, link.FriendId)
    if insertErr != nil {
        fmt.Println("Error when trying to save link")
        return err
    }
    fmt.Println("Link added")
    return nil
}
