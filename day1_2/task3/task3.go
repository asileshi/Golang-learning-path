package task3

import (
	"unicode"
)

func CheckPalindrome(s string) bool {

	i := 0
	j := len(s) - 1
	for i < j {
		if unicode.IsPunct(rune(s[i])) || unicode.IsSpace(rune(s[i])) {
			i++
			continue
		}
		if unicode.IsPunct(rune(s[j])) || unicode.IsSpace(rune(s[j])) {
			j--
			continue
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
