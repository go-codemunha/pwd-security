package validator

import (
	"fmt"
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
			{
				hasUppercase = true
				fmt.Printf("IsUpper %t : %q\n", hasUppercase, r)
			}
		case unicode.IsLower(r):
			{
				hasLowercase = true
				fmt.Printf("IsLower %t : %q\n", hasLowercase, r)
			}
		case unicode.IsNumber(r):
			{
				hasNumber = true
				fmt.Printf("IsNumber %t : %q\n", hasNumber, r)
			}
		case unicode.IsSpace(r):
			{
				hasSpace = false
				fmt.Printf("IsSpace %t : %q\n", hasSpace, r)
			}
		case !unicode.IsSymbol(r):
			{
				hasSymbol = true
				fmt.Printf("IsSymbol %t : %q\n", hasSymbol, r)
			}
		default:
		}

	}
	// fmt.Println(len(s), hasLowercase, hasUppercase, hasNumber, hasSpace, hasSymbol)
	return hasLowercase && hasUppercase && hasNumber && hasSpace && hasSymbol
}
