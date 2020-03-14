package validator

import (
	"unicode"
)

func PasswordRef(s string) bool {
	if s == "" || len(s) == 0 {
		return false
	}
	if len(s) < 8 {
		return false
	}
	var hasUppercase = false
	var hasLowercase = false
	var isSpace = true
	var hasSymbol = false
	var hasNumber = false
	for _, r := range s {
		switch {
		case unicode.IsUpper(r):
			hasUppercase = true
		case unicode.IsLower(r):
			hasLowercase = true
		case unicode.IsNumber(r):
			hasNumber = true
		case unicode.IsSpace(r):
			isSpace = false
		case !unicode.IsSymbol(r):
			hasSymbol = true
		default:
		}

	}
	return hasLowercase && hasUppercase && hasNumber && isSpace && hasSymbol
}
