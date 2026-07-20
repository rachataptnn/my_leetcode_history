package main

import "strconv"

func main() {
	summaryRanges(nil)
}

func summaryRanges(nums []int) []string {
	prev := nums[0]
	prevStr := strconv.Itoa(prev)
	if len(nums) == 1 {
		return []string{prevStr}
	}

	result := []string{}
	for i := 2; i < len(nums)-1; i++ {
		next := nums[i]
		if next - 
		str := gotTheRange(curr, next)
	}

	return nil
}

func IsContinue()
