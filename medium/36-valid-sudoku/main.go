// https://leetcode.com/problems/valid-sudoku/?envType=problem-list-v2&envId=hash-table&difficulty=MEDIUM

package main

import "fmt"

func main() {
	// board := [][]byte{
	// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	// board := [][]byte{{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	// 473/507
	board := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'}}

	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	rowInvalid := everyRowInvalid(board)
	colInvalid := everyColInvalid(board)
	subInvalid := every3x3Invalid(board)

	if rowInvalid || colInvalid || subInvalid {
		return false
	}

	return true
}

func everyRowInvalid(board [][]byte) bool {
	for _, v1 := range board {
		numberMap := make(map[byte]int)
		for _, v2 := range v1 {
			if v2 != '.' {
				numberMap[v2]++
				if numberMap[v2] > 1 {
					return true
				}
			}
		}
	}

	return false
}

func everyColInvalid(board [][]byte) bool {
	for col := 0; col < 9; col++ {
		numberMap := make(map[byte]int)
		for row := 0; row < 9; row++ {
			num := board[row][col]
			if num != '.' {
				numberMap[num]++
				if numberMap[num] > 1 {
					return true
				}
			}

		}
	}

	return false
}

func every3x3Invalid(board [][]byte) bool {
	arr012 := []int{0, 1, 2}
	arr345 := []int{3, 4, 5}
	arr678 := []int{6, 7, 8}

	for _, v1 := range arr012 {
		numberMap := make(map[byte]int)
		for _, v2 := range arr012 {
			num := board[v1][v2]
			if num != '.' {
				numberMap[num]++
				if numberMap[num] > 1 {
					return true
				}
			}
		}
	}
	for _, v1 := range arr012 {
		numberMap := make(map[byte]int)
		for _, v2 := range arr345 {
			num := board[v1][v2]
			if num != '.' {
				numberMap[num]++
				if numberMap[num] > 1 {
					return true
				}
			}
		}
	}
	for _, v1 := range arr012 {
		numberMap := make(map[byte]int)
		for _, v2 := range arr678 {
			num := board[v1][v2]
			if num != '.' {
				numberMap[num]++
				if numberMap[num] > 1 {
					return true
				}
			}
		}
	}

	for _, v1 := range arr345 {
		numberMap := make(map[byte]int)
		for _, v2 := range arr012 {
			num := board[v1][v2]
			if num != '.' {
				numberMap[num]++
				if numberMap[num] > 1 {
					return true
				}
			}
		}
	}
	for _, v1 := range arr345 {
		numberMap := make(map[byte]int)
		for _, v2 := range arr345 {
			num := board[v1][v2]
			if num != '.' {
				numberMap[num]++
				if numberMap[num] > 1 {
					return true
				}
			}
		}
	}
	for _, v1 := range arr345 {
		numberMap := make(map[byte]int)
		for _, v2 := range arr678 {
			num := board[v1][v2]
			if num != '.' {
				numberMap[num]++
				if numberMap[num] > 1 {
					return true
				}
			}
		}
	}

	for _, v1 := range arr678 {
		numberMap := make(map[byte]int)
		for _, v2 := range arr012 {
			num := board[v1][v2]
			if num != '.' {
				numberMap[num]++
				if numberMap[num] > 1 {
					return true
				}
			}
		}
	}
	for _, v1 := range arr678 {
		numberMap := make(map[byte]int)
		for _, v2 := range arr345 {
			num := board[v1][v2]
			if num != '.' {
				numberMap[num]++
				if numberMap[num] > 1 {
					return true
				}
			}
		}
	}
	for _, v1 := range arr678 {
		numberMap := make(map[byte]int)
		for _, v2 := range arr678 {
			num := board[v1][v2]
			if num != '.' {
				numberMap[num]++
				if numberMap[num] > 1 {
					return true
				}
			}
		}
	}

	return false
}
