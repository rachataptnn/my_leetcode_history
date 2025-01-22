package main

import "fmt"

func main() {
	// nums := []int{2, 1, 6, 4} // Output: 	1
	nums := []int{1, 1, 1} // Output: 	1
	fmt.Println(waysToMakeFair(nums))
}

func waysToMakeFair(nums []int) int {
	cnt := 0
	for i := range nums {
		sumOdd := 0
		sumEven := 0

		isSwap := false
		for j, v2 := range nums {
			if j == i {
				isSwap = true
				continue
			}

			if isSwap {
				if j%2 == 0 {
					sumOdd += v2
				} else {
					sumEven += v2
				}
			} else {
				if j%2 == 0 {
					sumEven += v2
				} else {
					sumOdd += v2
				}
			}

		}
		if sumOdd == sumEven {
			cnt++
		}
	}
	return cnt
}
