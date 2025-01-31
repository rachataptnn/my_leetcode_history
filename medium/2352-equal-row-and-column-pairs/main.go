package main

import "fmt"

func main() {
	grid := [][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7},
	}

	fmt.Println(equalPairs(grid))
}

func equalPairs(grid [][]int) int {
	n := len(grid)
	transpose := make([][]int, n)
	for i := range transpose {
		transpose[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j, rows := range grid {
			transpose[i][j] = rows[i]
		}
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isArrEq(transpose[i], grid[j]) {
				cnt++
			}
		}
	}

	return cnt
}

func isArrEq(arr1, arr2 []int) bool {
	for i, v := range arr1 {
		if arr2[i] != v {
			return false
		}
	}

	return true
}
