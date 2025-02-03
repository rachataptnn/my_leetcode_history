package main

import "fmt"

func main() {
	// case 1
	nums := []int{1, 4, 3, 3, 2}

	// case 3
	// nums := []int{3, 2, 1}

	// 465/868
	// nums := []int{1, 9, 7, 1}

	fmt.Println(longestMonotonicSubarray(nums))
}

func longestMonotonicSubarray(nums []int) int {
	longest := 1
	if len(nums) == 1 {
		return longest
	}

	inc := 1
	dec := 1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			dec++
			inc = 1
		}

		if nums[i] > nums[i+1] {
			inc++
			dec = 1
		}

		if dec > longest {
			longest = dec
		}
		if inc > longest {
			longest = inc
		}
		if nums[i] == nums[i+1] {
			inc, dec = 1, 1
		}
	}

	return longest
}
