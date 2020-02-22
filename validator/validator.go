package validator

import (
	"unicode"
)

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func Password(s string) bool {
	if s == "" || len(s) == 0 {
		return false
	}
	if len(s) < 8 {
		return false
	}
	var hasUppercase = false
	var hasLowercase = false
	var hasSpace = true
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
			hasSpace = false
		case !unicode.IsSymbol(r):
			hasSymbol = true
		default:
		}

	}
	return hasLowercase && hasUppercase && hasNumber && hasSpace && hasSymbol
}
