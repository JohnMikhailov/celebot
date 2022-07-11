package app

import "github.com/meehighlov/celebot/app/db"


func IsAuthUser(userId int) bool {
	user := db.User{ID: userId}
	isExist, err := user.IsExist()
	if err != nil {
		return false
	}
	return isExist
}

func IsAdmin(userId int) bool {
	user := db.User{ID: userId}
	err := user.Get()
	if err != nil {
		return false
	}

	return user.HasAdminAccess()
}
