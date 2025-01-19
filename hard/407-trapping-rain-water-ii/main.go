package main

import "fmt"

func main() {
	heightMap := [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}
	fmt.Println(trapRainWater(heightMap))
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) <= 1 {
		return 0
	}

	floors := wallsToFloors(heightMap)

	cols := len(heightMap)
	rows := len(heightMap[0])
	waterCells := (cols * rows) - 4

	sum := 0
	for i := len(floors) - 1; i > 0; i-- {
		area := findIslandArea(floors[i])
		if area <= waterCells {
			sum += area
		}
	}

	return sum
}

func wallsToFloors(walls [][]int) [][][]int {
	highestPillar := 0
	for _, wall := range walls {
		for _, pillar := range wall {
			if highestPillar < pillar {
				highestPillar = pillar
			}
		}
	}

	cols := len(walls)
	rows := len(walls[0])
	depth := highestPillar

	floors := make([][][]int, depth+1) // sometime i want index start with 1 lol
	for i := range floors {
		floors[i] = make([][]int, rows)
		for j := range floors[i] {
			floors[i][j] = make([]int, cols)
		}

	}

	for i, wall := range walls {
		for j, pillar := range wall {
			floors[pillar][j][i] = 1

			for index := pillar - 1; index > 0; index-- {
				floors[index][j][i] = 1
			}
		}
	}

	return floors
}

// helper func
func findIslandArea(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	totalArea := 0

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] != 0 {
			return 0
		}

		grid[row][col] = -1 // Mark as visited

		area := 1
		area += dfs(row-1, col)
		area += dfs(row+1, col)
		area += dfs(row, col-1)
		area += dfs(row, col+1)

		return area
	}

	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			if grid[row][col] == 0 {
				totalArea += dfs(row, col)
			}
		}
	}

	return totalArea
}
