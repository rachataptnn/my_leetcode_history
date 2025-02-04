package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 5, 10, 50}

	// nums := []int{10, 20, 30, 40, 50}

	fmt.Println(maxAscendingSum(nums))
}

func maxAscendingSum(nums []int) int {
	res, sum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		res = max(res, sum)
	}

	return res
}
