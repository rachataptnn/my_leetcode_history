// https://leetcode.com/problems/minimum-number-of-days-to-disconnect-island/description/?envType=daily-question&envId=2024-08-11

package main

import (
	"fmt"
)

type input struct {
	grid [][]int
}

func main() {
	// input := input{
	// 	grid: [][]int{
	// 		{0, 1, 1, 0},
	// 		{0, 1, 1, 0},
	// 		{0, 0, 0, 0}},
	// }

	// 2/99 holi shet -> actually its 88/99 lol
	// nid to change
	// if j+1 <= len(grid[0]) to
	// if j+1 < len(grid[0])
	// input := input{
	// 	grid: [][]int{
	// 		{1, 0, 1, 0},
	// 	},
	// }

	input := input{
		grid: [][]int{
			{1, 1, 0, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 0, 1, 1},
			{1, 1, 0, 1, 1},
		},
	}

	fmt.Println(minDays(input.grid))
}

func minDays(grid [][]int) int {
	// ive no idea bout handling this 'special' case
	if len(grid) == 1 {
		if len(grid[0]) == 2 {
			if grid[0][0] == 1 && grid[0][1] == 1 {
				return 2
			}
		}
	}

	min := 5

	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				cnt1S := getSurround1SCount(grid, i, j)
				if cnt1S < min {
					min = cnt1S
				}
			}
		}
	}

	return min
}

func getSurround1SCount(grid [][]int, i, j int) int {
	over := 0
	if i-1 >= 0 {
		over = grid[i-1][j]
	}

	left := 0
	if j-1 >= 0 {
		left = grid[i][j-1]
	}

	right := 0
	if j+1 < len(grid[0]) {
		right = grid[i][j+1]
	}

	below := 0
	if i+1 < len(grid) {
		below = grid[i+1][j]
	}

	return over + left + right + below
}
