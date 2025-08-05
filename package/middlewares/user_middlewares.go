package middlewares

import (
	"regexp"
	"unicode"
)

func Email_verification(email string) bool {
	emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
	re := regexp.MustCompile(emailRegex)
	if re.MatchString(email) {
		return true
	} else {
		return false
	}
}

func Password_verification(password string) bool {
	score := 0
	for _, char := range password {
		if unicode.IsUpper(char) {
			score = score + 1
		} else if unicode.IsLower(char) {
			continue
		} else if unicode.IsDigit(char) {
			score = score + 1
		} else {
			score = score + 1
		}
	}
	if score >= 5 {
		return true
	} else {
		return false
	}
}
