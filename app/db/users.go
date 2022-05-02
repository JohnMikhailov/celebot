package db


import "fmt"


func (user User) Save(userLinks []UserLink) error {
	stmt, err := Client.Prepare("INSERT INTO user(name, birthDate, link) VALUES($1, $2, $3) RETURNING id;")
    if err != nil {
        fmt.Println("Error when trying to prepare statement")
        fmt.Println(err)
        return err
    }
    defer stmt.Close()

	var userId userId

    insertErr := stmt.QueryRow(user.Name, user.BirthDate).Scan(&userId)
    if insertErr != nil {
        fmt.Println("Error when trying to save todo")
        return err
    }
    fmt.Println("User added")

	for _, userLink := range userLinks {
		userLink.Save()  // possible bottle neck - ".Save()" - is a db call
		user.UserLinks = append(user.UserLinks, userLink)
	}

    return nil
}

func (userLink UserLink) Save() error {
	stmt, err := Client.Prepare("INSERT INTO userLink(name, link, userId) VALUES($1, $2, $3);")
	if err != nil {
        fmt.Println("Error when trying to prepare statement")
        fmt.Println(err)
        return err
    }
    defer stmt.Close()
	insertErr := stmt.QueryRow(userLink.Name, userLink.Link, userLink.UserId)
    if insertErr != nil {
        fmt.Println("Error when trying to save todo")
        return err
    }
    fmt.Println("UserLink added")
    return nil
}
