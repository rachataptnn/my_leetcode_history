// https://leetcode.com/problems/trapping-rain-water/

package main

import (
	"fmt"
)

func main() {
	// height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1} // 6
	// height := []int{4, 2, 0, 3, 2, 5} // 9

	height := []int{0, 2, 0}

	fmt.Println(trap(height))
}

func trap(height []int) int {
	sum := 0

	sumWaterLR, ht := sumWaterLeftToRight(height)

	if ht.index < len(height)-2 {
		sumWaterRL := sumWaterRightToLeft(height, ht)
		for _, v := range sumWaterRL {
			sum += v
		}

		for i := 0; i < ht.index; i++ {
			sum += sumWaterLR[i]
		}
	} else {
		for _, v := range sumWaterLR {
			sum += v
		}
	}

	return sum
}

type highestTotem struct {
	height int
	index  int
}

func sumWaterLeftToRight(height []int) (sumWater []int, ht highestTotem) {
	max := 0
	for i, h := range height {
		if h > max {
			max = h
			ht.height = h
			ht.index = i
		}

		water := max - h
		sumWater = append(sumWater, water)
	}

	return sumWater, ht
}

func sumWaterRightToLeft(height []int, ht highestTotem) (sumWater []int) {
	max := 0
	for i := len(height) - 1; i >= ht.index; i-- {
		if height[i] > max {
			max = height[i]
		}

		water := max - height[i]
		sumWater = append(sumWater, water)
	}

	return sumWater
}
