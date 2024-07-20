// https://leetcode.com/problems/find-valid-matrix-given-row-and-column-sums/description/?envType=daily-question&envId=2024-07-20

package main

import "fmt"

func main() {
	// rowSum := []int{3, 8}
	// colSum := []int{4, 7}

	rowSum := []int{14, 9}
	colSum := []int{6, 9, 8}

	fmt.Println(restoreMatrix(rowSum, colSum))
}

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	rowAmt, colAmt := len(rowSum), len(colSum)
	res := prepareEmptyMatrix(rowAmt, colAmt)

	for i := 0; i < rowAmt; i++ {
		for j := 0; j < colAmt; j++ {
			res[i][j] = getMin(rowSum[i], colSum[j])
			rowSum[i] -= res[i][j]
			colSum[j] -= res[i][j]
		}
	}

	return res
}

func prepareEmptyMatrix(rowAmt, colAmt int) [][]int {
	matrix := make([][]int, rowAmt)
	for i := range matrix {
		matrix[i] = make([]int, colAmt)
	}
	return matrix
}

func getMin(rowVal, colVal int) int {
	if rowVal < colVal {
		return rowVal
	}

	return colVal
}
