package mark

import (
	"math/rand"

	"github.com/mattia01017/tictactia/board"
	. "github.com/mattia01017/tictactia/constant"
)

// CpuMove put a symbol for cpu
func CpuMove(b *board.Board) {
	// checks if cpu is about to win and make move
	l, c := oneLeft(*b, CPU_SYM_VAL)
	if l != ERR {
		(*b)[l][c] = CPU_SYM_VAL
		return
	}

	//checks if user is about to win and make a move
	l, c = oneLeft(*b, USER_SYM_VAL)
	if l != ERR {
		(*b)[l][c] = CPU_SYM_VAL
		return
	}

	// if nothing above, does a random move
	for {
		l = rand.Intn(3)
		c = rand.Intn(3)
		if (*b)[l][c] == 0 {
			(*b)[l][c] = CPU_SYM_VAL
			return
		}
	}
}

// oneLeft returns coordinates for the only empty cell left for winning if
// it exists, (-1,-1) in other cases
func oneLeft(b board.Board, player int) (l, c int) {
	//check lines
	for l = 0; l < 3; l++ {
		ml := make(map[int]int)
		for c = 0; c < 3; c++ {
			ml[b[l][c]]++
		}
		if ml[player] == 2 && ml[0] == 1 {
			// search the last empty in the selected line
			for c = 0; c < 3; c++ {
				if b[l][c] == 0 {
					return
				}
			}
		}
	}

	//check columns
	for c = 0; c < 3; c++ {
		mc := make(map[int]int)
		for l = 0; l < 3; l++ {
			mc[b[l][c]]++
		}
		if mc[player] == 2 && mc[0] == 1 {
			// search the last empty in the selected column
			for l = 0; l < 3; l++ {
				if b[l][c] == 0 {
					return
				}
			}
		}
	}

	//check diagonals
	md1 := make(map[int]int)
	md2 := make(map[int]int)
	for i := 0; i < 3; i++ {
		md1[b[i][i]]++
		md2[b[i][2-i]]++
	}
	if md1[player] == 2 && md1[0] == 1 {
		// search the last empty in the first diagonal
		for i := 0; i < 3; i++ {
			if b[i][i] == 0 {
				return i, i
			}
		}
	}
	if md2[player] == 2 && md2[0] == 1 {
		// search the last empty in the second diagonal
		for i := 0; i < 3; i++ {
			if b[i][2-i] == 0 {
				return i, 2 - i
			}
		}
	}
	return ERR, ERR
}
