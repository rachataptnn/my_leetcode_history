package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 5, 10, 50}
	fmt.Println(maxAscendingSum(nums))
}

func maxAscendingSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	max := 0
	tmpSum := 0
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 {
			if nums[i] < nums[i+1] {
				tmpSum += nums[i]
				if max < tmpSum {
					max = tmpSum
				}
			} else {
				tmpSum += nums[i]
				if max < tmpSum {
					max = tmpSum
				}
				tmpSum = 0
			}
		} else {
			if nums[i] > nums[i-1] {
				tmpSum += nums[i]
				if max < tmpSum {
					max = tmpSum
				}
			}
		}
	}

	return max
}
