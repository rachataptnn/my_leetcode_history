// https://leetcode.com/problems/build-a-matrix-with-conditions/description/?envType=daily-question&envId=2024-07-21

package main

import "fmt"

func main() {
	k := 3
	rowConditions := [][]int{{1, 2}, {3, 2}}
	colConditions := [][]int{{2, 1}, {3, 2}}

	buildMatrix(k, rowConditions, colConditions)
}

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	sortedRow := cutAndJoin(rowConditions)
	sortedCol := cutAndJoin(colConditions)

	fmt.Println(sortedRow, sortedCol)

	return nil
}

func cutAndJoin(matrix [][]int) []int {
	if len(matrix) == 1 {
		return matrix[0]
	}

	for len(matrix) > 1 {
		merged := mergeArr(matrix[0], matrix[1])
		matrix[0] = merged
		matrix = append(matrix[:1], matrix[1+1:]...) // remove 2nd element
		fmt.Println("")
	}

	return nil
}

func mergeArr(firstArr, secondArr []int) []int {
	for firstIndex, v := range firstArr {
		for secondIndex, val := range secondArr {
			if v == val {
				if firstIndex < secondIndex { // 2nd arr lead
					concatArr := append(secondArr[:secondIndex], firstArr[firstIndex:]...)
					return concatArr
				} else if firstIndex > secondIndex { // 2nd arr follow
					concatArr := append(firstArr[:firstIndex], secondArr[secondIndex:]...)
					return concatArr
				}
			}
		}
	}

	return firstArr
}
