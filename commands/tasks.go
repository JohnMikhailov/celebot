package commands

import (
	"fmt"
	"time"

	"github.com/meehighlov/celebot/app"
	"github.com/meehighlov/celebot/app/db"
	"github.com/meehighlov/celebot/telegram"
)

func CheckBirthDays(struct{}) {
	// TODO load data partialy - potentialy there could be a lot of data about birthdays even for 1 day
	// solution is: load (for example) 10 birthdays, send notifications, expose data about those 10 birthdays,
	// load another part of data

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
	// timeout := app.GetConfig().DEFAULT_DELAY_BETWEEN_REMINDINGS_SEC  // one day
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
