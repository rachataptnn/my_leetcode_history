// https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/description/?envType=problem-list-v2&envId=greedy

package main

import "fmt"

func main() {
	position := []int{1, 2, 3}
	// position := []int{2, 2, 2, 3, 3}
	// position := []int{1, 1000000000}
	fmt.Println(minCostToMoveChips(position))
}

func minCostToMoveChips(position []int) int {

	// 2, 2, 2, 3, 3

	// 2
	// 2 3
	// 2 3      -> 2 is highest freq, so it's pillar

	pillar := 0
	freqM := make(map[int]int)
	for _, v := range position {
		freqM[v]++
		if freqM[v] > pillar {
			pillar = v
		}
	}

	totalCost := 0
	isPillarOdd := pillar%2 != 0
	if isPillarOdd {
		for key := range freqM {
			if key == pillar {
				continue
			}

			if key%2 == 0 {
				totalCost += freqM[key]
			}
		}
	} else {
		for key := range freqM {
			if key == pillar {
				continue
			}

			if key%2 != 0 {
				totalCost += freqM[key]
			}
		}
	}

	return totalCost
}
