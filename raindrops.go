package raindrops

import "strconv"

func Convert(number int) string {
	var result string
	var hasFactors bool

	if number%3 == 0 {
		hasFactors = true
		result += "Pling"
	}

	if number%5 == 0 {
		hasFactors = true
		result += "Plang"
	}

	if number%7 == 0 {
		hasFactors = true
		result += "Plong"
	}

	if !hasFactors {
		return strconv.Itoa(number)
	} else {
		return result
	}
}
