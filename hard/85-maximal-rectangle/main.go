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
	// matrix := [][]byte{
	// 	{'0', '1'}} // Output: 1

	// 44/74
	// matrix := [][]byte{
	// 	{'1', '0', '1', '1', '1'},
	// 	{'0', '1', '0', '1', '0'},
	// 	{'1', '1', '0', '1', '1'},
	// 	{'1', '1', '0', '1', '1'},
	// 	{'0', '1', '1', '1', '1'}} // Output: 6

	// 69/74
	// matrix := [][]byte{
	// 	{'0', '1', '1', '0', '0', '1', '0', '1', '0', '1'},
	// 	{'0', '0', '1', '0', '1', '0', '1', '0', '1', '0'},
	// 	{'1', '0', '0', '0', '0', '1', '0', '1', '1', '0'},
	// 	{'0', '1', '1', '1', '1', '1', '1', '0', '1', '0'},
	// 	{'0', '0', '1', '1', '1', '1', '1', '1', '1', '0'},
	// 	{'1', '1', '0', '1', '0', '1', '1', '1', '1', '0'},
	// 	{'0', '0', '0', '1', '1', '0', '0', '0', '1', '0'},
	// 	{'1', '1', '0', '1', '1', '0', '0', '1', '1', '1'},
	// 	{'0', '1', '0', '1', '1', '0', '1', '0', '1', '1'}} // output: 10

	// 65/74
	// matrix := [][]byte{
	// 	{'1', '1', '1'},
	// 	{'0', '0', '0'},
	// 	{'0', '1', '0'},
	// 	{'0', '1', '0'},
	// 	{'0', '1', '0'},
	// 	{'0', '0', '1'},
	// 	{'1', '0', '0'}} // output 3

	// 70/74
	matrix := [][]byte{
		{'1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '0', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'0', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1'},
		{'1', '0', '0', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '0', '0', '0', '1', '1', '1', '1', '1', '0', '1', '0'},
		{'1', '0', '1', '1', '0', '0', '0', '1', '1', '1', '1', '0', '1', '0', '1'},
		{'1', '0', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1'},
		{'1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '0', '0', '0', '1', '0', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '0', '1'}}

	fmt.Println("MAX : ", maximalRectangle(matrix))
}

//TODO: calculatedMap[startFlr:endFlr::startCol:endCol]area

type building struct {
	floors []floor
}

type floor struct {
	lines []line
}

type line struct {
	start int
	end   int
}

func maximalRectangle(matrix [][]byte) int {
	building := building{}

	maxArea := 0
	for _, row := range matrix {
		floor := floor{}
		lines := getContinuousLines(row)
		floor.lines = append(floor.lines, lines...)
		building.floors = append(building.floors, floor)
	}

	for i, floor := range building.floors {
		fmt.Printf("\n\n\n\nFLOOR: %d\n-------", i)
		for _, line := range floor.lines {
			fmt.Printf("\nline: %d, %d\n", line.start, line.end)
			area := getAreaByLine(line, i, building)
			if area > maxArea {
				maxArea = area
			}

			areaS := getAreaByShrinkableLine(line, i, building)
			fmt.Printf("Area-: %d, Area-s: %d", area, areaS)
			if areaS > maxArea {
				maxArea = areaS
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

func getAreaByLine(line line, floorNumber int, building building) (upperLines int) {
	possibleFloorCnt := 1
	explorationBlocked := true

	// explore upper floors
	currentFloor := floorNumber - 1
	for currentFloor >= 0 {
		upperFloor := building.floors[currentFloor]
		explorationBlocked = true
		for _, upperLine := range upperFloor.lines {
			if upperLine.start <= line.start && upperLine.end >= line.end {
				possibleFloorCnt++
				explorationBlocked = false
			}
		}
		if explorationBlocked {
			break
		}
		currentFloor--
	}

	// explore lower floors
	if floorNumber == 3 {
		fmt.Println(3)
	}
	currentFloor = floorNumber + 1
	for currentFloor < len(building.floors) {
		lowerFloor := building.floors[currentFloor]
		explorationBlocked = true
		for _, lowerLine := range lowerFloor.lines {
			if lowerLine.start <= line.start && lowerLine.end >= line.end {
				possibleFloorCnt++
				explorationBlocked = false
			}
		}
		if explorationBlocked {
			break
		}
		currentFloor++
	}

	lineLength := line.end - line.start + 1
	area := possibleFloorCnt * lineLength
	return area
}

func getAreaByShrinkableLine(line line, floorNumber int, building building) (upperLines int) {
	if floorNumber == 3 {
		fmt.Println("3")
	}

	possibleFloorCnt := 1
	explorationBlocked := true

	// start, end := 0, 0

	// explore upper floors
	currentFloor := floorNumber - 1
	for currentFloor >= 0 {
		upperFloor := building.floors[currentFloor]
		explorationBlocked = len(upperFloor.lines) == 0
		for _, upperLine := range upperFloor.lines {
			if isLinesOverlap(upperLine, line) && !explorationBlocked {
				line.start = max(upperLine.start, line.start)
				line.end = min(upperLine.end, line.end)
				possibleFloorCnt++
				explorationBlocked = false
			} else {
				explorationBlocked = true
			}
		}
		if explorationBlocked || len(upperFloor.lines) == 0 {
			break
		}
		currentFloor--
	}

	// explore lower floors
	currentFloor = floorNumber + 1
	for currentFloor < len(building.floors) {
		lowerFloor := building.floors[currentFloor]
		explorationBlocked = len(lowerFloor.lines) == 0
		for _, lowerLine := range lowerFloor.lines {
			if isLinesOverlap(lowerLine, line) && !explorationBlocked {
				line.start = max(lowerLine.start, line.start)
				line.end = min(lowerLine.end, line.end)
				possibleFloorCnt++
				explorationBlocked = false
			} else {
				explorationBlocked = true
			}
		}
		if explorationBlocked {
			break
		}
		currentFloor++
	}

	lineLength := line.end - line.start + 1
	area := possibleFloorCnt * lineLength
	return area
}

func recursiveUpper(inputLine line, startFloor int, building building) (stackedFloor int) {
	currentFloor := startFloor - 1
	if currentFloor <= 0 {
		return stackedFloor
	}

	upperFloor := building.floors[currentFloor]
	for _, line := range upperFloor.lines {
		if isLinesOverlap(inputLine, line) {
			inputLine.start = max(inputLine.start, line.start)
			inputLine.end = min(inputLine.end, line.end)
		}
	}

	return recursiveUpper(inputLine, currentFloor, building)
}

func recursiveLineShorter()

func isLinesOverlap(line1, line2 line) bool {
	return line1.start <= line2.end && line1.end >= line2.start
}
