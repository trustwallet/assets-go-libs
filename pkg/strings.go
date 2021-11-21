package pkg

import (
	"strings"
	"unicode"
)

func Contains(str string, entries []string) bool {
	for _, f := range entries {
		if f == str {
			return true
		}
	}

	return false
}

func IsLowerCase(str string) bool {
	return str == strings.ToLower(str)
}

func IsUpperCase(str string) bool {
	return str == strings.ToUpper(str)
}

func ReverseCase(str string) string {
	n := 0
	runes := make([]rune, len(str))
	for _, r := range str {
		runes[n] = r
		n++
	}
	runes = runes[0:n]

	for i := 0; i < n; i++ {
		if unicode.IsLower(runes[i]) {
			runes[i] = unicode.ToUpper(runes[i])
		} else if unicode.IsUpper(runes[i]) {
			runes[i] = unicode.ToLower(runes[i])
		} else {
			runes[i] = runes[i]
		}
	}

	return string(runes)
}
