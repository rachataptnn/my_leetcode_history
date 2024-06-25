// https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips/description/?envType=daily-question&envId=2024-06-24

package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	nums := []int{0, 0, 0, 1, 0, 1, 1, 0}
	k := 3

	// nums := []int{0, 1, 0}
	// k := 1

	result := minKBitFlips(nums, k)
	fmt.Println(result)

	endTime := time.Now()
	fmt.Println(startTime.Sub(endTime))
}

func minKBitFlips(nums []int, k int) int {
	flipCnt := 0
	for {
		index := findFirstZero(nums)
		if index == -1 {
			return flipCnt
		}

		subSlice := getSubSlice(nums, k, index)
		if subSlice == nil {
			return -1
		}

		flip(subSlice)
		flipCnt++
	}
}

func findFirstZero(arr []int) int {
	for i, val := range arr {
		if val == 0 {
			return i
		}
	}

	return -1
}

func getSubSlice(arr []int, k int, index int) []int {
	if index+k > len(arr) {
		return nil
	}
	return arr[index : index+k]
}

func flip(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] ^= 1
	}
}
