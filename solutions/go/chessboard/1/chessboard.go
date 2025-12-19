package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f, found := cb[file]
    if !found {
        return 0
    }
    
    occupied := 0
    for _, value := range f {
        if value == true {
            occupied++
        }
    }

    return occupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
        return 0
    }
    
    occupied := 0
    for key, _ := range cb {
        row := cb[key]
        for index, _ := range row {
            if index == rank - 1 && row[index] == true {
                occupied++
            }
        }
    }

    return occupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
    for key, _ := range cb {
        for range cb[key] {
            count++
        } 
    }
    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
    for key, _ := range cb {
        for _, v := range cb[key] {
            if v == true {
                count++
            }
        } 
    }
    return count
}
