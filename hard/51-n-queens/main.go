// https://leetcode.com/problems/n-queens/description/

package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 4
	fmt.Println(solveNQueens(n))
}

// i need to learn backtracking! look like this problem must be solve in O(n!)
func solveNQueens(n int) [][]string {
	switch n {
	case 1:
		return [][]string{{"Q"}}
	case 2:
		return nil
	case 3:
		return nil
	}

	return generateBoards(n)
}

func generateBoards(n int) [][]string {
	templateBoard := genTemplateBoard(n)

	possibleBoards := [][]string{}
	for row := 0; row < len(templateBoard); row++ {
		copiedBoard := make([][]string, len(templateBoard))
		for i := range templateBoard {
			copiedBoard[i] = make([]string, len(templateBoard[i]))
			copy(copiedBoard[i], templateBoard[i])
		}

		board := genBoard(copiedBoard, row, n)
		if board != nil {
			possibleBoards = append(possibleBoards, board)
		}
	}

	return possibleBoards
}

func genTemplateBoard(n int) [][]string {
	board := [][]string{}
	for i := 0; i < n; i++ {
		row := make([]string, n)
		for j := 0; j < n; j++ {
			row[j] = "."
		}
		board = append(board, row)
	}
	return board
}

func genBoard(board [][]string, row, n int) []string {
	// orginalRow := row
	col := 0

	queenAmt := 0
	for col < len(board) {
		if !isQueenSeenQueen(board, row, col) {
			board[row][col] = "Q"
			queenAmt++
			col++

			row = 0
			// row++
			// if row >= len(board[0]) {
			// 	row = 0
			// 	fmt.Println("")
			// }
		} else {
			row++
			if row >= len(board[0]) {
				col++
				row = 0
				fmt.Println("")
			}
		}
	}

	if queenAmt == n {
		res := []string{}
		for _, v := range board {
			row := strings.Join(v, "")
			res = append(res, row)
		}

		return res
	}

	return nil
}

func isQueenSeenQueen(board [][]string, row, col int) bool {
	isSeen := isQSQVertical(board, col) || isQSQHorizontal(board, row) || isQSQDiagonal(board, row, col)
	return isSeen
}

func isQSQVertical(board [][]string, col int) bool {
	for i := range board {
		if board[i][col] == "Q" {
			return true
		}
	}

	return false
}

func isQSQHorizontal(board [][]string, row int) bool {
	bRow := board[row]
	for _, v := range bRow {
		if v == "Q" {
			return true
		}
	}

	return false
}

func isQSQDiagonal(board [][]string, row, col int) bool {
	originRow, originCol := row, col

	// to top left -1 -1
	// break when: row < 0 || col < 0
	for {
		row--
		col--
		if row < 0 || col < 0 {
			break // check dia to top left, didn't see any Queen!, lets check other Diagonal
		}

		if board[row][col] == "Q" {
			return true
		}
	}
	row, col = originRow, originCol

	// to bottom left +1 -1
	// break when: row >= len(board[0]) || col < 0
	for {
		row++
		col--
		if row >= len(board[0]) || col < 0 {
			break // check dia to bottom left, didn't see any Queen!, lets check other Diagonal
		}

		if board[row][col] == "Q" {
			return true
		}
	}
	row, col = originRow, originCol

	// to top right -1 +1
	// break when: row < 0 || col >= len(board)
	for {
		row--
		col++
		if row < 0 || col >= len(board) {
			break // check dia to top right, didn't see any Queen!, lets check other Diagonal
		}

		if board[row][col] == "Q" {
			return true
		}
	}
	row, col = originRow, originCol

	// to bottom right +1 +1
	// break when: row >= len(board[0]) || col >= len(board)
	for {
		row++
		col++
		if row >= len(board[0]) || col >= len(board) {
			break // check dia to bottom right, didn't see any Queen!, DIDNOTSEEANYQUEEN!!!
		}

		if board[row][col] == "Q" {
			return true
		}
	}

	return false // DIDNOTSEEANYQUEEN!!!
}
