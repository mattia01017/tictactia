package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/mattia01017/tictactia/board"
	. "github.com/mattia01017/tictactia/constant"
	"github.com/mattia01017/tictactia/mark"
)

const DELAY = 1000000000 // in nanoseconds

func main() {
	flush()
	rand.Seed(time.Now().UnixNano())
	var b board.Board
	var cin, lin int

	fmt.Printf("Welcome!\n%v\n", b)

	// game starts
	for {

		// player turn
		for {
			fmt.Print("Insert a coordinate (line,column): ")
			fmt.Scanf("%d,%d", &lin, &cin)

			if mark.PlayerMove(lin, cin, &b) {
				flush()
				fmt.Printf("\n%v\n", b)
				break
			} else {
				fmt.Println("Input isn't valid. Retry")
			}
		}

		// check if the player won
		if board.Win(b, USER_SYM_VAL) {
			fmt.Printf(USER_COLOR, "You won!\n")
			os.Exit(0)
			// check if board is full
		} else if board.IsFull(b) {
			fmt.Println("No one is the winner...")
			os.Exit(0)
		}

		// CPU turn
		fmt.Print("CPU is moving... \n\n")
		time.Sleep(DELAY)
		flush()
		mark.CpuMove(&b)
		fmt.Printf("\n%v\n", b)

		// check if cpu won
		if board.Win(b, CPU_SYM_VAL) {
			fmt.Printf(CPU_COLOR, "You lost!\n")
			os.Exit(0)
			// check if board is full
		} else if board.IsFull(b) {
			fmt.Println("No one is the winner...")
			os.Exit(0)
		}
	}
}

// clear the terminal
func flush() {
	out, _ := exec.Command("clear").Output()
	fmt.Print(string(out))
}
