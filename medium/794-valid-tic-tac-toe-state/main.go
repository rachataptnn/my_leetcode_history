package main

import "fmt"

func main() {
	board := []string{"XOX", "O O", "XOX"}
	fmt.Println(validTicTacToe(board))
}

func validTicTacToe(board []string) bool {
	xCnt := 0
	oCnt := 0

	for _, row := range board {
		for _, cell := range row {
			if cell == 'X' {
				xCnt++
			}
			if cell == 'O' {
				oCnt++
			}
		}
	}

	turnIncorrect := xCnt-oCnt > 1
	xLessThanO := oCnt > xCnt
	if turnIncorrect || xLessThanO {
		return false
	}

	sus := isSusWin(board)

	return !sus
}

func isSusWin(board []string) bool {
	xWin := 0
	oWin := 0

	firstCol := ""
	secCol := ""
	thirdCol := ""
	for i, row := range board {
		if row == "XXX" {
			xWin++
		} else if row == "OOO" {
			oWin++
		}

		if i == 0 {
			firstCol += string(row[i])
		} else if i == 1 {
			secCol += string(row[i])
		} else {
			thirdCol += string(row[i])
		}
	}

	if firstCol == "XXX" || secCol == "XXX" || thirdCol == "XXX" {
		xWin++
	}
	if firstCol == "OOO" || secCol == "OOO" || thirdCol == "OOO" {
		oWin++
	}

	susWin := xWin > 1 || oWin > 1 || (xWin > 0 && oWin > 0)

	return susWin
}
