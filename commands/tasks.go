package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/meehighlov/celebot/app"
	"github.com/meehighlov/celebot/app/db"
	"github.com/meehighlov/celebot/telegram"
)

func getUsersToNitificate(dayWithMonth string, limit, offset int) []db.Friend {
	user := db.User{}
	user.GetFriendsByBirthDate(dayWithMonth, limit, offset)
	log.Println("friends found:", len(user.Friends), "for day:", dayWithMonth)

	return user.Friends
}

func CheckBirthDays(struct{}) {
	log.Println("searching for birthdays")

	config := app.GetConfig()
	client := telegram.NewApiClient(config.BOTTOKEN_CELEBOT)
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
			text := fmt.Sprintf("Today is a birthday of %s!", user.Name)
			client.SendMessage(user.GetChatIdStr(), text, false)
		}

		offset += shift
	}
}

func RunChecks() {
	timeout := app.GetConfig().DEFAULT_DELAY_BETWEEN_CHECKS_SEC
	tasksQueue := make(chan struct{}, 1)
	hourToNotify := app.GetConfig().BD_NOTIFICATION_HOUR_MOSCOW_TZ
	location, _ := time.LoadLocation("Europe/Moscow")

	go func() {
		for v := range tasksQueue {
			CheckBirthDays(v)
		}
	}()
	go func() {
		for {
			now := time.Now().In(location).Hour()
			if now == hourToNotify {
				tasksQueue <- struct{}{}
				hour := 60 * 60
				time.Sleep(time.Duration(hour) * time.Second)
			}
			time.Sleep(timeout)
		}
	}()
}
