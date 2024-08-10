// https://leetcode.com/problems/stone-game-ix/

package main

import "fmt"

func main() {
	stones := []int{2, 1}
	fmt.Println(stoneGameIX(stones))
}

func stoneGameIX(stones []int) bool {
	round := pickNum(stones, 0, 1)
	return round%2 == 0
}

func pickNum(stones []int, sum, round int) int {
	selectedNum := 0
	selectedIdx := 0
	for i, v := range stones {
		if sum+v%3 != 0 {
			selectedIdx = i
			selectedNum = v

			sum += v
			round++

			break
		}
	}

	if selectedNum == 0 && selectedIdx == 0 { // end case
		return round
	}

	stones = removeElementByIndex(stones, selectedIdx)
	return pickNum(stones, sum, round)
}

func removeElementByIndex(arr []int, index int) []int {
	if index >= 0 && index < len(arr) {
		return append(arr[:index], arr[index+1:]...)
	}
	return arr
}
