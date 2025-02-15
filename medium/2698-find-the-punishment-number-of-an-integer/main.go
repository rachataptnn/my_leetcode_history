package main

import (
	"fmt"
)

func main() {
	fmt.Println(punishmentNumber(37))
}

func punishmentNumber(n int) int {

	punishmentNum := 0

	for currentNum := 1; currentNum <= n; currentNum++ {
		squareNum := currentNum * currentNum

		// Check if valid partition can be found and add squared number if so
		if canPartition(squareNum, currentNum) {
			punishmentNum += squareNum
		}
	}

	return punishmentNum
}

func canPartition(num, target int) bool {
	if target < 0 || num < target {
		return false
	}

	if num == target {
		return true
	}

	return canPartition(num/10, target-(num%10)) ||
		canPartition(num/100, target-(num%100)) ||
		canPartition(num/1000, target-(num%1000))
}
