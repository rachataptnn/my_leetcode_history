package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(check(nums))
}

func check(nums []int) bool {
	n := len(nums)

	circularNums := append(nums, nums...)
	fmt.Println(circularNums, n)

	for i := 0; i < n; i++ {
		sliced := circularNums[i : i+n]
		if isSorted(sliced) {
			return true
		}
	}

	return false
}

func isSorted(arr []int) bool {
	prev := 0
	for _, v := range arr {
		if v < prev {
			return false
		}
		prev = v
	}

	return true
}
