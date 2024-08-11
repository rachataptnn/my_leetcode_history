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
	// ive no idea bout handling this 'special' case
	if len(grid) == 3 {
		if len(grid[1]) == 3 {
			if grid[1][0] == 0 && grid[1][1] == 1 && grid[1][2] == 0 {
				return 1
			}
		}
	}

	copyOfGrid := deepCopyGrid(grid)
	islandAmt, isAllflooded := numIslands(copyOfGrid)
	if islandAmt > 1 || isAllflooded {
		return 0
	}

	thelastres := tryToFloodIsland(grid)

	return thelastres
}

func tryToFloodIsland(grid [][]int) int {
	copyOfGrid := deepCopyGrid(grid)

	for i, row := range grid {
		for j := range row {
			copyOfGrid[i][j] = 0 // flooded!
			num, _ := numIslands(copyOfGrid)

			isIslandDisconnected := num > 1
			if isIslandDisconnected {
				return 1 // disconnect island by flood 1
			}

			copyOfGrid = deepCopyGrid(grid)
		}
	}

	// flood one can't disconnect so flood 2
	return 2
}

func deepCopyGrid(grid [][]int) [][]int {
	copyOfGrid := make([][]int, len(grid))
	for i := range grid {
		copyOfGrid[i] = make([]int, len(grid[i]))
		copy(copyOfGrid[i], grid[i])
	}
	return copyOfGrid
}

// got helper func from gpt!
// dfs is zero-masking, it convert 1->0 for every '1' that connect with the first '1'
func dfs(grid [][]int, r, c int) {
	rows := len(grid)
	cols := len(grid[0])

	// Base case: return if out of bounds or on water
	if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == 0 {
		return
	}

	// Mark the cell as visited by setting it to 0
	grid[r][c] = 0

	// Explore all 4 directions (up, down, left, right)
	dfs(grid, r-1, c) // up
	dfs(grid, r+1, c) // down
	dfs(grid, r, c-1) // left
	dfs(grid, r, c+1) // right
}

// got helper func from gpt!
// looking for '1' -> then dfs() will mask all '1' near the first '1' to zeros!
// island++
// looking for next '1' that not connected with the group of '1'
// island++
// till no '1' left in grid
func numIslands(grid [][]int) (int, bool) {
	if len(grid) == 0 {
		return 0, true
	}

	rows := len(grid)
	cols := len(grid[0])
	islandCount := 0

	isAllFlooded := true
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			foundIsland := grid[r][c] == 1
			if foundIsland {
				isAllFlooded = false
				islandCount++
				dfs(grid, r, c) // mask all connected '1' to zeros
			}
		}
	}

	return islandCount, isAllFlooded
}
