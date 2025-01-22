package main

import "fmt"

func main() {
	// nums := []int{2, 1, 6, 4} // Output: 	1
	nums := []int{1, 1, 1} // Output: 	1
	fmt.Println(waysToMakeFair(nums))
}

func waysToMakeFair(nums []int) int {
	sumOdd := 0
	sumEven := 0

	fairCnt := 0

	for i, v := range nums {
		if i%2 == 0 {
			sumOdd += v
		} else {
			sumEven += v
		}
	}

	for i, v := range nums {
		if i%2 == 0 {
			tmpSumEven := sumEven - v
			if tmpSumEven == sumOdd {
				fairCnt++
			}
		} else {
			tmpSumOdd := sumOdd - v
			if tmpSumOdd == sumEven {
				fairCnt++
			}
		}
	}

	return fairCnt
}
