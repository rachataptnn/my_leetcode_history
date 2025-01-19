package main

import (
	"fmt"
	"sort"
)

func main() {
	// stones := []int{2, 7, 4, 1, 8, 1}

	stones := []int{9, 3, 2, 10}

	fmt.Println(lastStoneWeight(stones))
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 2 {
		sort.Ints(stones)
		return stones[1] - stones[0]
	}

	stone := recFn(stones)

	return stone[0]
}

func recFn(stones []int) []int {
	if len(stones) == 1 {
		return stones
	}
	if len(stones) == 0 {
		return []int{0}
	}

	sort.Ints(stones)
	biggestStone := stones[len(stones)-1]
	bigStone := stones[len(stones)-2]
	removedStones := stones[:len(stones)-2]

	newStone := biggestStone - bigStone
	if newStone > 0 {
		removedStones = append(removedStones, newStone)
	}

	return recFn(removedStones)
}
