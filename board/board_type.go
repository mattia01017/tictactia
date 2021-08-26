package board

import (
	"fmt"

	. "github.com/mattia01017/tictactia/constant"
)

// Board is a 3x3 table. The board[n][m] value is 0 if empty, 1 if filled by the user, 2 if filled by cpu
// Board[line][column]
type Board [3][3]int

/*
aspect:
    1   2   3
1 | x | o | x |
  -------------
2 | o | o | o |
  -------------
3 | x | x | o |
*/

// String method is to manage the way the board is displayed by fmt functions
func (b Board) String() (out string) {
	out += "    1   2   3\n"
	for l := 0; l < 5; l++ {

		if l%2 == 1 {
			out += "  -------------\n"
		} else {
			line := l / 2
			out += fmt.Sprintf("%d ", line+1)

			for c := 0; c < 3; c++ {
				switch b[line][c] {
				case 0:
					out += "|   "
				case CPU_SYM_VAL:
					out += fmt.Sprintf("| %s ", fmt.Sprintf(CPU_COLOR, CPU_SYM))
				case USER_SYM_VAL:
					out += fmt.Sprintf("| %s ", fmt.Sprintf(USER_COLOR, USER_SYM))
				}
			}
			out += "|\n"
		}
	}
	return
}
