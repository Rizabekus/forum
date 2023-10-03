package pkg

import (
	"regexp"
)

func PostChecker(title string, text string) (bool, string) {
	if len(title) < 2 {
		return false, "Too short for title."
	} else if len(text) < 2 {
		return false, "Too short for post."
	}
	if len(title) > 40 {
		return false, "Too long for title."
	}

	return true, "OK"
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
