package strand

func ToRNA(dna string) string {
	mapping := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

	var rnaString string

	for _, ch := range dna {
		rnaString += string(mapping[ch])
	}

	return rnaString
}
