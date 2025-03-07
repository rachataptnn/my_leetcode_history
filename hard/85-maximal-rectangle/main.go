package main

import "fmt"

func main() {
	// matrix := [][]byte{
	// 	{'1', '0', '1', '0', '0'},
	// 	{'1', '0', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1'},
	// 	{'1', '0', '0', '1', '0'}} // Output: 6

	// matrix := [][]byte{
	// 	{'1'}} // Output: 1

	// 42/74
	matrix := [][]byte{
		{'0', '1'}} // Output: 1

	fmt.Println(maximalRectangle(matrix))
}

type line struct {
	start int
	end   int
}

func maximalRectangle(matrix [][]byte) int {
	maxArea := 0

	for i, v := range matrix {
		lines := getContinuousLines(v)

		for _, line := range lines {
			areaByThisLine := getAreaByExploreUpperAndLower(matrix, line.start, line.end, i)
			fmt.Println("line:", line.start, line.end)
			fmt.Println("Area by this line:", areaByThisLine)

			if maxArea < areaByThisLine {
				maxArea = areaByThisLine
			}
		}
	}

	return maxArea
}

func getContinuousLines(floor []byte) (lines []line) {
	start, end := 0, 0
	foundLine := false

	for i, v := range floor {
		isStartLine := v == '1' && !foundLine
		isEndLine := v == '0' && foundLine

		isMostRightStartLine := v == '1' && i == len(floor)-1 && !foundLine
		isMostRightEndLine := v == '1' && i == len(floor)-1 && foundLine

		if isMostRightStartLine {
			start, end = i, i
			lines = append(lines, line{start, end})
			foundLine = false
		} else if isMostRightEndLine {
			end = i
			lines = append(lines, line{start, end})
			foundLine = false
		} else if isStartLine {
			start = i
			foundLine = true
		} else if isEndLine {
			end = i - 1
			lines = append(lines, line{start, end})
			foundLine = false
		}
	}

	return lines
}

func getAreaByExploreUpperAndLower(matrix [][]byte, start, end, currentFloor int) int {
	// upper
	floor := currentFloor - 1
	goodFloorCount := 0
	for floor >= 0 {
		upperFloor := matrix[floor]
		hasHole := false
		for i := start; i <= end; i++ {
			if upperFloor[i] == '0' {
				hasHole = true
				break
			}
		}

		if !hasHole {
			goodFloorCount++
		}
		floor--
	}

	// lower
	floor = currentFloor + 1
	for floor < len(matrix) {
		lowerFloor := matrix[floor]
		hasHole := false
		for i := start; i <= end; i++ {
			if lowerFloor[i] == '0' {
				hasHole = true
				break
			}
		}

		if !hasHole {
			goodFloorCount++
		}
		floor++
	}

	areaByThisLine := (end - start + 1) * (goodFloorCount + 1)
	return areaByThisLine
}
