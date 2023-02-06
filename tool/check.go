package tool

import (
	"strings"
	"suning/dao"
)

func IsShortName(username string) bool {
	var name = []rune(username)
	if len(name) > 0 {
		return true
	} else {
		return false
	}
}

func IsNameErr(username string) bool {
	if strings.Contains(username, "\\") == true {
		return false
	} else {
		return true
	}
}

func IsRepeatUsername(username string) bool {
	U := dao.TakeUserDate()
	var usernames []string
	ok := false
	for _, date := range U {
		usernames = append(usernames, date.Username)
	}
	for _, value := range usernames {
		if value == username {
			ok = true
		} else {
			continue
		}
	}
	if ok {
		return ok
	} else {
		return !ok
	}
}

func IsPasswordErr(password string) bool {
	if strings.Contains(password, "\\") == true {
		return false
	} else {
		return true
	}
}

func IsShortPassword(password string) bool {
	var name = []rune(password)
	if len(name) >= 6 {
		return true
	} else {
		return false
	}
}

func IsLongPassword(password string) bool {
	var name = []rune(password)
	if len(name) <= 12 {
		return true
	} else {
		return false
	}
}
