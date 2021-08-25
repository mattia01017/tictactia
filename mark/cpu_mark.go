package mark

import (
	"math/rand"

	"github.com/mattia01017/tictactia/board"
	. "github.com/mattia01017/tictactia/constant"
)

// CpuMove put a symbol for cpu
func CpuMove(b *board.Board) {
	// check if cpu is about to win and make move
	l, c := oneLeft(*b, CPU_SYM_VAL)
	if l != ERR {
		(*b)[l][c] = CPU_SYM_VAL
		return
	}

	//check if user is about to win and make a move
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

// returns coordinates for the only one left
func oneLeft(b board.Board, who int) (l, c int) {
	//check lines
	for l = 0; l < 3; l++ {
		ml := make(map[int]int)
		for c = 0; c < 3; c++ {
			ml[b[l][c]]++
		}
		if ml[who] == 2 && ml[0] == 1 {
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
		if mc[who] == 2 && mc[0] == 1 {
			for l = 0; l < 3; l++ {
				if b[l][c] == 0 {
					return
				}
			}
		}
	}

	//check diagonal 1
	md1 := make(map[int]int)
	md2 := make(map[int]int)
	for i := 0; i < 3; i++ {
		md1[b[i][i]]++
		md2[b[i][2-i]]++
	}
	if md1[who] == 2 && md1[0] == 1 {
		for i := 0; i < 3; i++ {
			if b[i][i] == 0 {
				return i, i
			}
		}
	}
	if md2[who] == 2 && md2[0] == 1 {
		for i := 0; i < 3; i++ {
			if b[i][2-i] == 0 {
				return i, 2 - i
			}
		}
	}
	return ERR, ERR
}
