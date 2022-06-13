package app

import "log"

func IsAllowedUser(username string) bool {
	if !GetConfig().IsUsernameExist(username) {
		log.Println("Not permited user " + username + " tried to call celebot")
		return false
	}
	return true
}
