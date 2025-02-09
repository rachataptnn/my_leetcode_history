package main

import "fmt"

func main() {
	nums := []int{4, 1, 3, 3}
	fmt.Println(countBadPairs(nums))
}

func countBadPairs(nums []int) int64 {
	badPairCnt := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if i < j && (j-i) != nums[j]-nums[i] {
				badPairCnt++
			}
		}
	}

	return int64(badPairCnt)
}
