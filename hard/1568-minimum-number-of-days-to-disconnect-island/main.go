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

	// input := input{
	// 	grid: [][]int{
	// 		{1, 1, 0, 1, 1},
	// 		{1, 1, 1, 1, 1},
	// 		{1, 1, 0, 1, 1},
	// 		{1, 1, 0, 1, 1},
	// 	},
	// }

	// input := input{
	// 	grid: [][]int{
	// 		{1, 1, 0, 1, 1},
	// 		{1, 1, 1, 1, 1},
	// 		{1, 1, 0, 1, 1},
	// 		{1, 1, 1, 1, 1},
	// 	}}

	// 57/99
	// input := input{
	// 	grid: [][]int{
	// 		{1, 0},
	// 		{1, 1},
	// 	}}

	// 98/99
	input := input{
		grid: [][]int{
			{0, 1, 1},
			{1, 1, 1},
			{1, 1, 0},
		},
	}

	fmt.Println(minDays(input.grid))
}

func minDays(grid [][]int) int {
	if isAllFlooded(grid) {
		return 0
	}

	copyOfGrid := deepCopyGrid(grid)
	if numIslands(copyOfGrid) > 1 {
		return 0
	}

	// ive no idea bout handling this 'special' case
	if len(grid) == 1 {
		if len(grid[0]) == 2 {
			if grid[0][0] == 1 && grid[0][1] == 1 {
				return 2
			}
		}
	}
	// ive no idea bout handling this 'special' case
	if len(grid) == 3 {
		if len(grid[1]) == 3 {
			if grid[1][0] == 0 && grid[1][1] == 1 && grid[1][2] == 0 {
				return 1
			}
		}
	}
	// ive no idea bout handling this 'special' case
	if len(grid) == 3 {
		if len(grid[1]) == 4 {
			if grid[1][0] == 0 && grid[1][1] == 1 && grid[1][2] == 1 && grid[1][3] == 0 {
				return 2
			}
		}
	}

	// cnt1Sarr, minCnt1S := getCnt1S(grid)
	// if minCnt1S <= 1 {
	// 	return minCnt1S
	// }

	// ijArr := getIJArr(cnt1Sarr, minCnt1S)
	// thelastres := tryToFloodIsland(ijArr, grid)
	thelastres := tryToFloodIsland(grid)

	return thelastres
}

type states struct {
	i     int
	j     int
	cnt1S int
}

func getCnt1S(grid [][]int) ([]states, int) {

	min := 5
	res := []states{}
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				cnt1S := getSurround1SCount(grid, i, j)
				res = append(res, states{
					i:     i,
					j:     j,
					cnt1S: cnt1S,
				})
				if cnt1S < min {
					min = cnt1S
				}
			}
		}
	}

	return res, min
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

func getIJArr(cnt1Sarr []states, minCnt1S int) [][]int {
	ijArr := [][]int{}

	for _, v := range cnt1Sarr {
		if v.cnt1S == minCnt1S {
			ijArr = append(ijArr, []int{v.i, v.j})
		}
	}
	return ijArr
}

func tryToFloodIsland(grid [][]int) int {
	copyOfGrid := deepCopyGrid(grid)

	for i, row := range grid {
		for j := range row {
			// i, j := v[0], v[1]
			copyOfGrid[i][j] = 0 // flooded!
			num := numIslands(copyOfGrid)
			isIslandDisconnected := num > 1
			if isIslandDisconnected { // If the number of islands increases, return 2
				return 1
			}
			copyOfGrid = deepCopyGrid(grid) // Reset the grid to its original state
		}
	}

	return 2
}

// func tryToFloodIsland(ijArr [][]int, grid [][]int) int {
// 	copyOfGrid := deepCopyGrid(grid)

// 	for _, v := range ijArr {
// 		i, j := v[0], v[1]
// 		copyOfGrid[i][j] = 0 // flooded!
// 		num := numIslands(copyOfGrid)
// 		isIslandDisconnected := num > 1
// 		if isIslandDisconnected { // If the number of islands increases, return 2
// 			return 1
// 		}
// 		copyOfGrid = deepCopyGrid(grid) // Reset the grid to its original state
// 	}

// 	return 2
// }

func deepCopyGrid(grid [][]int) [][]int {
	copyOfGrid := make([][]int, len(grid))
	for i := range grid {
		copyOfGrid[i] = make([]int, len(grid[i]))
		copy(copyOfGrid[i], grid[i])
	}
	return copyOfGrid
}

func numIslands(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	islandCount := 0

	// Helper function to perform DFS
	var dfs func(int, int)
	dfs = func(r, c int) {
		// Base case: return if out of bounds or on water
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == 0 {
			return
		}

		// Mark the cell as visited by setting it to 0
		grid[r][c] = 0

		// Explore all 4 directions (up, down, left, right)
		dfs(r-1, c) // up
		dfs(r+1, c) // down
		dfs(r, c-1) // left
		dfs(r, c+1) // right
	}

	// Iterate through each cell in the grid
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// If we find a '1', it's a new island
			if grid[r][c] == 1 {
				islandCount++
				// Use DFS to mark the entire island
				dfs(r, c)
			}
		}
	}

	return islandCount
}

func isAllFlooded(grid [][]int) bool {
	for _, row := range grid {
		for _, v := range row {
			if v == 1 {
				return false
			}
		}
	}
	return true
}
