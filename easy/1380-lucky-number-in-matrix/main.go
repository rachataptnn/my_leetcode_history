// https: //leetcode.com/problems/lucky-numbers-in-a-matrix/?envType=daily-question&envId=2024-07-19

package main

import "fmt"

func main() {
	// input := [][]int{
	// 	{3, 7, 8},
	// 	{9, 11, 13},
	// 	{15, 16, 17}}

	input := [][]int{
		{7, 8},
		{1, 2},
	}
	fmt.Println(luckyNumbers(input))
}

func luckyNumbers(matrix [][]int) []int {
	for i := 0; i < len(matrix); i++ {
		min := matrix[i][0]
		minIndex := 0
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
				minIndex = j
			}
		}

		max := 0
		maxIndex := 0
		for k := 0; k < len(matrix); k++ {
			if matrix[k][minIndex] > max {
				max = matrix[k][minIndex]
				maxIndex = k
			}
		}

		if maxIndex == i {
			return []int{matrix[maxIndex][minIndex]}
		}
	}
	return nil
}
