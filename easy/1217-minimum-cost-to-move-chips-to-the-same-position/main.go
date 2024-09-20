// https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/description/?envType=problem-list-v2&envId=greedy

package main

import "fmt"

func main() {
	// position := []int{1, 2, 3}
	// position := []int{2, 2, 2, 3, 3}
	// position := []int{1, 1000000000}

	// 40/51
	// position := []int{2, 3, 3}

	// 39/51
	position := []int{6, 4, 7, 8, 2, 10, 2, 7, 9, 7} // <- this is not and Final Visualization, need to transform it to Frequency map
	// Frequency map
	//   |   |   |   |   |   | o |   |   |
	//   | o |   |   |   |   | o |   |   |
	//   | o |   | o |   | o | o | o | o | o
	// 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10

	// and then just try to seek
	// 1. "highest pillar of even" (2nd)
	// 1. "highest pillar of odd" (7th)

	// if we pick the even pillar,
	// - all the even coins is does'nt have any cost to move
	// - so we focus only odd coins (7th, 9th)
	// next, we move any odd coin to even pillar (2nd)
	// cost is 4 (3 from 7th, and 1 from 9th)

	// do it again by picking odd pillar
	// 6 (2 from 2nd, 1 from 4th, 1 from 6th, 1 from 8th, and 1 from 10th)

	// so the minimum cost is: 4

	fmt.Println(minCostToMoveChips(position))
}

func minCostToMoveChips(position []int) int {
	oddCoins, evenCoins := 0, 0
	for _, v := range position {
		if v%2 != 0 {
			oddCoins++
		} else {
			evenCoins++
		}
	}

	if oddCoins < evenCoins {
		return oddCoins
	}

	return evenCoins
}
