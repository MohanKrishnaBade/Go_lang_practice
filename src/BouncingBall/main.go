package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	_"time"
)

func main() {

	const (
		width     = 50
		height    = 10
		cellBall  = '%'
		cellEmpty = ' '
		maxFrames = 1200
	)

	var (
		cell   rune
		vx, vy = 1, 1
		px, py int
	)

	screen.Clear()
	//this will creates an muti dimenshional array.
	board := make([][]bool, width)
	for index := range board {
		board[index] = make([]bool, height)
	}

	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		board[px][py] = true
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				fmt.Print(string(cell))
			}
			screen.MoveTopLeft()
			//  time.Sleep(time.Second /10)
		}
	}
	// fmt.Println(board)
}
