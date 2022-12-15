package chessboard

// Rank stores if a square is occupied by a piece - this will be a slice of bools.
type Rank []bool

// Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	var result int
	for _, s := range cb[rank] {
		if s {
			result++
		}
	}
	return result
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	var result int
	if file < 1 || file > 8 {
		return result
	}
	for _, rank := range cb {
		if rank[file-1] {
			result++
		}
	}
	return result
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb["A"])
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var result int
	for rank := range cb {
		result += CountInRank(cb, rank)
	}
	return result
}
