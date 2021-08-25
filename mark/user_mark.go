package mark

import "github.com/mattia01017/tictactia/board"

// player move puts a user mark in the coordinate given. Returns true
// if the mark can be done
func PlayerMove(l, c int, b *board.Board) bool {
	if l > 3 || l < 1 || c > 3 || c < 1 {
		return false
	}
	if (*b)[l-1][c-1] != 0 {
		return false
	}
	(*b)[l-1][c-1] = 1
	return true
}
