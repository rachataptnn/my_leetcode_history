// https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips/description/?envType=daily-question&envId=2024-06-24

package main

import (
	"errors"
	"fmt"
)

func main() {
	nums := []int{0, 0, 0, 1, 0, 1, 1, 0}
	// nums := []int{0, 1, 0}
	k := 3

	result := minKBitFlips(nums, k)
	fmt.Println(result)
}

func minKBitFlips(nums []int, k int) int {
	flipCnt := 0

	for {
		index := findFirstZero(nums)
		if index == -1 {
			return flipCnt
		}

		subSlice, err := getSubSlice(nums, k, index)
		if err != nil {
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

func getSubSlice(arr []int, k int, index int) ([]int, error) {
	if index < 0 || index >= len(arr) || index+k > len(arr) {
		return nil, errors.New("invalid index or length")
	}
	return arr[index : index+k], nil
}

func flip(arr []int) {
	for i := 0; i < len(arr); i++ {
		switch arr[i] {
		case 0:
			arr[i] = 1
		case 1:
			arr[i] = 0
		default:
			fmt.Println("err: num is not 0,1")
			arr[i] = -1
		}
	}
}
