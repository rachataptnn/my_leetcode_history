package main

import "fmt"

func main() {
	// arr := []int{2, 8, 7, 4, 1, 3, 5, 6, 9}
	// mat := [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}

	// 460 / 1058 testcases passed
	arr := []int{1, 4, 5, 2, 6, 3}
	mat := [][]int{{4, 3, 5}, {1, 2, 6}}

	// 1055/1058
	// some really long long arrs
	fmt.Println(firstCompleteIndex(arr, mat))
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	ijCell := initMap(mat)

	colGoal := len(mat)
	rowGoal := len(mat[0])

	iMap := make(map[int]int)
	jMap := make(map[int]int)

	for ii, v := range arr {
		ij := ijCell[v]
		i := ij[0]
		j := ij[1]

		iMap[i]++
		if iMap[i] >= rowGoal {
			return ii
		}

		jMap[j]++
		if jMap[j] >= colGoal {
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
