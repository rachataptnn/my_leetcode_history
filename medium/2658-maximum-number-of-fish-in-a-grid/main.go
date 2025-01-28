package main

import "fmt"

func main() {
	grid := [][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}

	fmt.Println(findMaxFish(grid))
}

func findMaxFish(grid [][]int) int {
	rowAmt, colAmt := len(grid), len(grid[0])
	max := 0

	for r, row := range grid {
		for c := range row {
			// if cell != 0 {
			catchedFish := fishing(r, c, rowAmt, colAmt, grid)
			if catchedFish > max {
				max = catchedFish
			}
			// }
		}
	}

	return max
}

func fishing(r, c, rowAmt, colAmt int, grid [][]int) int {
	topFish, btmFish, lftFish, rgtFish := 0, 0, 0, 0

	if r-1 >= 0 {
		topFish += grid[r-1][c]
	}
	if r+1 < rowAmt {
		btmFish += grid[r+1][c]
	}
	if c-1 >= 0 {
		lftFish += grid[r][c-1]
	}
	if c+1 < colAmt {
		lftFish += grid[r][c+1]
	}

	return topFish + btmFish + lftFish + rgtFish
}
