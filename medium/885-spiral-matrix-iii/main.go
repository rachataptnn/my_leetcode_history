// https://leetcode.com/problems/spiral-matrix-iii/?envType=daily-question&envId=2024-08-08

package main

import "fmt"

func main() {
	// Input:
	rows := 1
	cols := 4
	rStart := 0
	cStart := 0

	// Output: [[0,0],[0,1],[0,2],[0,3]]
	fmt.Println(spiralMatrixIII(rows, cols, rStart, cStart))
}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	res := [][]int{}
	tmp := []int{}

	for i := rStart; i < rows; i++ {
		tmp = append(tmp, i)
		for j := cStart; j < cols; j++ {
			tmp = append(tmp, j)
			res = append(res, tmp)

			tmp = []int{}
			tmp = append(tmp, i)
			fmt.Println("")

		}
	}

	return res
}
