// https://leetcode.com/problems/magic-squares-in-grid/?envType=daily-question&envId=2024-08-09

package main

import "fmt"

func main() {
	// grid := [][]int{
	// 	{4, 3, 8, 4},
	// 	{9, 5, 1, 9},
	// 	{2, 7, 6, 2}}

	grid := [][]int{{8}}

	fmt.Println(numMagicSquaresInside(grid))
}

func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 {
		return 0
	}
	if len(grid[0]) < 3 {
		return 0
	}

	res := 0
	curRow := 0
	curCol := 0

	isDone := false

	for !isDone {
		row1 := []int{grid[curRow][curCol], grid[curRow][curCol+1], grid[curRow][curCol+2]}
		row2 := []int{grid[curRow+1][curCol], grid[curRow+1][curCol+1], grid[curRow+1][curCol+2]}
		row3 := []int{grid[curRow+2][curCol], grid[curRow+2][curCol+1], grid[curRow+2][curCol+2]}

		square := [][]int{row1, row2, row3}
		if isDistinct(square) {
			if isMagicSquare(square) {
				res++
			}
		}

		curCol++
		isVerticalEnd := curCol+2 == len(grid[0])
		if isVerticalEnd {
			curCol = 0
			curRow++

			isHorizontalEnd := curRow+2 == len(grid)
			if isHorizontalEnd {
				isDone = true
			}
		}
	}

	return res
}

func isDistinct(square [][]int) bool {
	seen := make(map[int]struct{})

	for _, row := range square {
		for _, num := range row {
			if _, exists := seen[num]; exists {
				return false // Duplicate found
			}
			seen[num] = struct{}{}
		}
	}

	return true
}

func isMagicSquare(square [][]int) bool {
	sumRow1 := square[0][0] + square[0][1] + square[0][2]
	sumRow2 := square[1][0] + square[1][1] + square[1][2]
	sumRow3 := square[2][0] + square[2][1] + square[2][2]

	sumCol1 := square[0][0] + square[1][0] + square[2][0]
	sumCol2 := square[0][1] + square[1][1] + square[2][1]
	sumCol3 := square[0][2] + square[1][2] + square[2][2]

	sumDiaLR := square[0][0] + square[1][1] + square[2][2]
	sumDiaRL := square[0][2] + square[1][1] + square[2][0]

	return sumRow1 == sumRow2 && sumRow2 == sumRow3 && // rows equal
		sumCol1 == sumCol2 && sumCol2 == sumCol3 && // cols equal
		sumDiaLR == sumDiaRL && // dias equal
		sumRow1 == sumCol1 && // rows eq cols
		sumRow1 == sumDiaLR // rows eq dias
}
