package main

import "fmt"

func main() {
	// nums := []int{1, 3, 2, 1, 5, 4}
	// k := 2

	nums := []int{2, 1}
	k := 1

	fmt.Println(sumOfGoodNumbers(nums, k))
}

func sumOfGoodNumbers(nums []int, k int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		if isGreaterThanLeft(nums, i, k) && isGreaterThanRight(nums, i, k) {
			sum += nums[i]
		}
	}

	return sum
}

func isGreaterThanLeft(nums []int, i, k int) bool {
	if i-k >= 0 {
		if nums[i] > nums[i-k] {
			return true
		}
	} else {
		return true
	}

	return false
}

func isGreaterThanRight(nums []int, i, k int) bool {
	if i+k < len(nums) {
		if nums[i] > nums[i+k] {
			return true
		}
	} else {
		return true
	}

	return false
}
