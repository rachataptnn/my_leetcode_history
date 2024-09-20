// https://leetcode.com/problems/can-place-flowers/description/?envType=problem-list-v2&envId=greedy&difficulty=EASY

package main

import "fmt"

func main() {
	// flowerbed := []int{1, 0, 0, 0, 1}

	// 105/130         0  1  2  3  4  5
	flowerbed := []int{1, 0, 0, 0, 0, 1}

	// 119/130
	// flowerbed := []int{1, 0, 0, 0, 1, 0, 0}
	n := 2

	//                 0  1  2  3  4  5  6  7  8  9  0  1  2  3
	// flowerbad := []int{1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1}
	// expect 0cnt = 5

	fmt.Println(canPlaceFlowers(flowerbed, n))
}

// i can actually o(n) with one loop
// add 0 to the start and the end of array flowerbed
// then loop start at index 1 end at index-2
// then check 3 condtions
// if isPrevious == 1 || isCurrent ==1 || isNext == 1 {
//   plantAble++
//   i++
// }

func canPlaceFlowers(flowerbed []int, n int) bool {
	unAvailableBed := countUnavailableBed(flowerbed)

	plantAble := 0
	for i := 0; i < len(flowerbed); i++ {
		if !unAvailableBed[i] {
			plantAble++
			i++
			// fmt.Println("we can plant on:", i)
		}
	}

	return plantAble >= n
}

func countUnavailableBed(flowerbed []int) map[int]bool {
	unAvailableBed := make(map[int]bool)
	for currentBed := 0; currentBed < len(flowerbed); currentBed++ {
		if flowerbed[currentBed] == 1 {
			previousBed := currentBed - 1
			unAvailableBed[previousBed] = true

			unAvailableBed[currentBed] = true

			nextBed := currentBed + 1
			unAvailableBed[nextBed] = true
		}
	}

	return unAvailableBed
}
