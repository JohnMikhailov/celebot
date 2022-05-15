package commands

import (
	"fmt"
	"time"

	"github.com/meehighlov/celebot/app"
	"github.com/meehighlov/celebot/app/db"
	"github.com/meehighlov/celebot/telegram"
)

func CheckBirthDays(struct{}) {
	fmt.Println("start task")
	user := db.User{}
	config := app.GetConfig()
	client := telegram.NewApiClient(config.BOTTOKEN)
	now := time.Now()
	dbDateFormat := now.Format("02.01.2006")
	dayWithMonth := dbDateFormat[:5]
	user.GetFriendsByBirthDate(dayWithMonth)

	fmt.Println("friends found:", len(user.Friends), "for day:", dayWithMonth)

	for _, friend := range user.Friends {
		text := fmt.Sprintf("birth date %s", friend.Name)
		client.SendMessage(friend.GetChatIdStr(), text)
	}
}

func RunChecks() {
	// timeout := time.Sleep(60 * 24 * time.Minute)  // one day
	tasksQueue := make(chan struct{}, 1)

	go func() {
		for v := range tasksQueue {
			CheckBirthDays(v)
		}
	}()
	go func() {
		for {
			time.Sleep(10 * time.Second)
			tasksQueue <-struct{}{}
		}
	}()
	// close(tasksQueue)
}
