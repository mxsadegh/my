package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const n = 12

var board [n]int

func vazir(x, y int) bool {

	for i := x - 1; i >= 0; i-- {

		if board[i] > 0 && (float64(x-i) == math.Abs(float64(board[i]-y)) || board[i] == y) {
			return false
		}
	}
	return true
}

var game [][]string

func gameboard() {

	for i, _ := range board {
		board[i] = 0
	}
}

func show() {

	game = make([][]string, n)
	for i := range game {
		game[i] = make([]string, n)
		for j := range game[i] {
			game[i][j] = "_"
		}
	}

	for i, j := range board {
		if j != 0 {
			game[i][j-1] = "Q"
		}

	}
	for i := 0; i < n; i++ {
		//fmt.Println(strings.Join(game[i], "|"), i)
		fmt.Fprintln(os.Stdout, strings.Join(game[i], "|"), i)
	}
	//fmt.Println("------------------------")
	fmt.Fprintln(os.Stdout, "--------------------------")
}

func main() {
	total := 0
	j := 1
	for i := 0; i < n; {

		for ; j <= n; j++ {

			if vazir(i, j) && board[i] != j {
				board[i] = j
				i++
				j = 1

				if i == n {
					show()
					i--
					total++
					j = board[i]
					board[i] = 0

					continue
				}

				break
			}

		}
		if j == n+1 {
			if i == 0 {
				break
			}
			board[i] = 0
			i--
			j = board[i]

		}

	}
	fmt.Println(total)
}
