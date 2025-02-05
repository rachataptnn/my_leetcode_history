package main

import "fmt"

func main() {
	// board := []string{"XOX", "O O", "XOX"}

	// 105/111
	// board := []string{"XO ", "XO ", "XO "}

	// 105/111
	// board := []string{
	// 	"XXX",
	// 	"OOX",
	// 	"OOX"}

	// 107/111
	// board := []string{
	// 	"XXX",
	// 	"XOO",
	// 	"OO "}

	// 108/111
	board := []string{
		"OXX",
		"XOX",
		"OXO"}

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

	xWin, oWin := isSusWin(board)
	susWin := xWin > 2 || oWin > 2 || (xWin > 0 && oWin > 0)
	if susWin {
		return false
	}

	if xWin == 1 && (oCnt >= xCnt) {
		return false
	}

	return true
}

func isSusWin(board []string) (int, int) {
	xWin := 0
	oWin := 0

	firstCol := ""
	secCol := ""
	thirdCol := ""
	for _, row := range board {
		if row == "XXX" {
			xWin++
		} else if row == "OOO" {
			oWin++
		}

		firstCol += string(row[0])
		secCol += string(row[1])
		thirdCol += string(row[2])
	}

	if firstCol == "XXX" || secCol == "XXX" || thirdCol == "XXX" {
		xWin++
	}
	if firstCol == "OOO" || secCol == "OOO" || thirdCol == "OOO" {
		oWin++
	}

	return xWin, oWin
}
