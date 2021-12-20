package strings

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
		switch {
		case unicode.IsLower(runes[i]):
			runes[i] = unicode.ToUpper(runes[i])
		case unicode.IsUpper(runes[i]):
			runes[i] = unicode.ToLower(runes[i])
		}
	}

	return string(runes)
}

func Difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}

	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}

	return diff
}
