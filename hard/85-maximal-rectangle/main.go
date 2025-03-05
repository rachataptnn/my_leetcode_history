package main

import "fmt"

func main() {
	matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}} // Output: 6
	fmt.Println(maximalRectangle(matrix))
}

func maximalRectangle(matrix [][]byte) int {
	for floor, curFloor := range matrix {
		isLastFloor := floor == len(matrix)-1
		if isLastFloor {
			break
		}

		lowerFlrLen, holes := surveyLowerFloor(matrix[floor+1], len(curFloor))
		if holes {
			continue
		}
	}

	return 0
}

func surveyCurrentFloor(currFlr []byte) int {
	longestLen, startIdx, endIdx := 0

	for i, v := range currFlr {

	}

	return longestLen, startIdx, endIdx
}

func surveyLowerFloor(lowerFlr []byte, curFlrLen int) (lowerFlrLen int, holes bool) {
	for len, tile := range lowerFlr {
		if tile == 0 {
			return len, len < curFlrLen
		}
	}

	return len(lowerFlr), false
}
