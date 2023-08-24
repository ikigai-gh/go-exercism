package isogram

import "strings"

func IsIsogram(word string) bool {
	m := make(map[rune]int)
	for _, c := range strings.ToLower(word) {
		if string(c) == " " || string(c) == "-" {
			continue
		}
		if m[c] == 0 {
			m[c] = 1
		} else {
			return false
		}
	}
	return true
}
