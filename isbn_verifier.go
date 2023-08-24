package isbn

import (
	"regexp"
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbnRegex := regexp.MustCompile(`^\d{9}[X,\d]$`)
	isbn = strings.Replace(isbn, "-", "", -1)

	if !isbnRegex.MatchString(isbn) {
		return false
	}

	var checkSum int

	for idx, ch := range isbn {
		if ch == 'X' {
			checkSum += 10 * (10 - idx)
		}
		num, _ := strconv.Atoi(string(ch))
		checkSum += num * (10 - idx)
	}

	return checkSum%11 == 0
}
