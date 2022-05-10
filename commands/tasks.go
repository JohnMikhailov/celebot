package commands

import (
	"fmt"
	"time"
	"github.com/meehighlov/celebot/app"
	"github.com/meehighlov/celebot/app/db"
	"github.com/meehighlov/celebot/telegram"
)

func CheckBirthDays() {
	user := db.User{}
	config := app.GetConfig()
	client := telegram.NewApiClient(config.BOTTOKEN)
	for {
		now := time.Now()
		dbDateFormat := now.Format("02.01.2006")
		dayWithMonth := dbDateFormat[:6]
		user.GetFriendsByBirthDate(dayWithMonth)

		for _, friend := range user.Friends {
			text := fmt.Sprintf("birth date %s", friend.Name)
			client.SendMessage(friend.GetChatIdStr(), text)
		}

		oneDay := 60 * 24 * time.Minute
		time.Sleep(oneDay)
	}
}
