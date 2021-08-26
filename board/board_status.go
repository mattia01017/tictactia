package board

// isFull returns true only if there's no empty cell left in the board
func IsFull(b Board) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[j][i] == 0 {
				return false
			}
		}
	}
	return true
}

// Win returns true only if there's a tris for the input player in the board
func Win(b Board, player int) bool {
	// check lines
	for l := 0; l < 3; l++ {
		ml := make(map[int]int)
		for c := 0; c < 3; c++ {
			ml[b[l][c]]++
		}
		if ml[player] == 3 {
			return true
		}
	}

	// check columns
	for c := 0; c < 3; c++ {
		mc := make(map[int]int)
		for l := 0; l < 3; l++ {
			mc[b[l][c]]++
		}
		if mc[player] == 3 {
			return true
		}
	}

	// check diagonals
	md1 := make(map[int]int)
	md2 := make(map[int]int)
	for i := 0; i < 3; i++ {
		md1[b[i][i]]++
		md2[b[i][2-i]]++
	}
	if md1[player] == 3 || md2[player] == 3 {
		return true
	}
	return false
}
