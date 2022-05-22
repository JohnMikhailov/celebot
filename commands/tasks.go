package commands

import (
	"fmt"
	"time"

	"github.com/meehighlov/celebot/app"
	"github.com/meehighlov/celebot/app/db"
	"github.com/meehighlov/celebot/telegram"
)


func getUsersToNitificate(dayWithMonth string, limit, offset int) []db.Friend {
	user := db.User{}
	user.GetFriendsByBirthDate(dayWithMonth, limit, offset)
	fmt.Println("friends found:", len(user.Friends), "for day:", dayWithMonth)

	return user.Friends
}

func CheckBirthDays(struct{}) {
	fmt.Println("start task")

	config := app.GetConfig()
	client := telegram.NewApiClient(config.BOTTOKEN)
	now := time.Now()
	dbDateFormat := now.Format("02.01.2006")
	dayWithMonth := dbDateFormat[:5]

	limit, offset := 10, 0
	shift := 10

	for {
		users := getUsersToNitificate(dayWithMonth, limit, offset)
		if len(users) == 0 {
			break
		}

		for _, user := range users {
			text := fmt.Sprintf("birth date %s", user.Name)
			client.SendMessage(user.GetChatIdStr(), text)
		}

		offset += shift
	}
}

func RunChecks() {
	timeout := app.GetConfig().DEFAULT_DELAY_BETWEEN_CHECKS_SEC
	tasksQueue := make(chan struct{}, 1)

	go func() {
		for v := range tasksQueue {
			CheckBirthDays(v)
		}
	}()
	go func() {
		for {
			time.Sleep(timeout)
			tasksQueue <-struct{}{}
		}
	}()
}
