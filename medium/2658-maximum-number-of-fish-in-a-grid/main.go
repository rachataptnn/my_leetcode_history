package main

import "fmt"

func main() {
	grid := [][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}

	fmt.Println(findMaxFish(grid))
}

func findMaxFish(grid [][]int) int {
	var (
		dfs      func(int, int) int
		maxCatch int
	)

	dfs = func(row, col int) int {
		if row < 0 || col < 0 || row == len(grid) || col == len(grid[0]) || grid[row][col] == 0 {
			return 0
		}

		fish := grid[row][col]
		grid[row][col] = 0
		return fish + dfs(row, col+1) + dfs(row+1, col) + dfs(row-1, col) + dfs(row, col-1)
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			maxCatch = max(maxCatch, dfs(row, col))
		}
	}
	return maxCatch
}
