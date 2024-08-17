// https://leetcode.com/problems/maximum-number-of-points-with-cost/?envType=daily-question&envId=2024-08-17

package main

import "fmt"

func main() {
	points := [][]int{
		{1, 2, 3},
		{1, 5, 1},
		{3, 1, 1},
	}
	fmt.Println(maxPoints(points))
}

func maxPoints(points [][]int) int64 {
	max := 0
	// rotated := points
	rotated := transpose(points)

	for i1 := 0; i1 < len(rotated); i1++ {
		src := rotated[i1][0]
		for i2 := 1; i2 < len(rotated); i2++ {
			sum := src
			for j := 0; j < len(rotated[0]); j++ {
				dst := rotated[i2][j]
				diff := abs(i1 - i2)
				sum += dst - diff
				if sum > max {
					max = sum
					fmt.Println("")
				}
			}
		}
	}

	return int64(max)
}

func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}

	numRows := len(matrix)
	numCols := len(matrix[0])
	transposed := make([][]int, numCols)

	for i := range transposed {
		transposed[i] = make([]int, numRows)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
