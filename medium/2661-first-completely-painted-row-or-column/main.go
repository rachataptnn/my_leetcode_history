package main

import "fmt"

func main() {
	// arr := []int{2, 8, 7, 4, 1, 3, 5, 6, 9}
	// mat := [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}

	// 460 / 1058 testcases passed
	arr := []int{1, 4, 5, 2, 6, 3}
	mat := [][]int{{4, 3, 5}, {1, 2, 6}}

	fmt.Println(firstCompleteIndex(arr, mat))
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	ijCell := initMap(mat)

	rowGoal := len(mat)
	colGoal := len(mat[0])
	for ii, v := range arr {
		ij := ijCell[v]
		i := ij[0]
		j := ij[1]

		mat[i][j] = -1
		if isCompletelyPainted(mat, i, j, rowGoal, colGoal) {
			return ii
		}
	}

	return 0
}

func initMap(mat [][]int) map[int][]int {
	ijCell := make(map[int][]int)

	for i, row := range mat {
		for j, cell := range row {
			ijCell[cell] = []int{i, j}
		}
	}

	return ijCell
}

func isCompletelyPainted(mat [][]int, i, j, rowGoal, colGoal int) bool {
	rowPainted := 0
	colPainted := 0
	for ii, row := range mat {
		for jj, cell := range row {
			if cell == -1 && ii == i {
				rowPainted++
				if rowPainted == rowGoal {
					return true
				}
			}

			if cell == -1 && jj == j {
				colPainted++
				if colPainted == colGoal {
					return true
				}
			}
		}
	}

	return false
}
