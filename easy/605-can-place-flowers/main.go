// https://leetcode.com/problems/can-place-flowers/description/?envType=problem-list-v2&envId=greedy&difficulty=EASY

package main

import "fmt"

func main() {
	// flowerbed := []int{1, 0, 0, 0, 1}

	// 105/130
	flowerbed := []int{1, 0, 0, 0, 0, 1}
	n := 2

	//                 0  1  2  3  4  5  6  7  8  9  0  1  2  3
	// flowerbad := []int{1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1}
	// expect 0cnt = 5

	fmt.Println(canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	unAvailableBed := countUnavailableBed(flowerbed)
	availableBed := len(flowerbed) - len(unAvailableBed)
	if availableBed%2 != 0 && availableBed != 1 {
		availableBed = availableBed + 1
	}

	return availableBed >= n
}

func countUnavailableBed(flowerbed []int) map[int]bool {
	unAvailableBed := make(map[int]bool)
	if flowerbed[0] == 1 {
		unAvailableBed[0] = true
		unAvailableBed[1] = true
	}
	if flowerbed[len(flowerbed)-1] == 1 {
		if len(flowerbed)-2 >= 0 {
			unAvailableBed[len(flowerbed)-2] = true
		}
		unAvailableBed[len(flowerbed)-1] = true
	}

	for currentBed := 1; currentBed < len(flowerbed)-2; currentBed++ {
		if flowerbed[currentBed] == 1 {
			previousBed := currentBed - 1
			unAvailableBed[previousBed] = true
			unAvailableBed[currentBed] = true

			if currentBed+1 < len(flowerbed) {
				nextBed := currentBed + 1
				unAvailableBed[nextBed] = true
			}
		}
	}

	return unAvailableBed
}
