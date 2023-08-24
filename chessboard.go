package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool

type Chessboard map[string]File

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	occupuiedCount := 0

	column, ok := cb[file]

	if !ok {
		return 0
	}

	for _, cellIsOccupuied := range column {
		if cellIsOccupuied {
			occupuiedCount++
		}
	}

	return occupuiedCount
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	occupiedCount := 0

	for _, rankN := range cb {
		if rankN[rank-1] {
			occupiedCount++
		}
	}

	return occupiedCount
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	cellCount := 0

	for _, c := range cb {
		for _, _ = range c {
			cellCount++
		}
	}

	return cellCount
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupiedCount := 0

	for _, rank := range cb {
		for _, cellIsOccupied := range rank {
			if cellIsOccupied {
				occupiedCount++
			}
		}
	}

	return occupiedCount
}
