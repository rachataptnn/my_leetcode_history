package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := []int{4, 14, 2}
	fmt.Println(totalHammingDistance(nums))
}

func totalHammingDistance(nums []int) int {
	totalHd := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		bin1 := IntToBinary(nums[i])
		for j := i + 1; j < n; j++ {
			if j == i {
				continue
			}

			bin2 := IntToBinary(nums[j])
			hd := HammingDistance(bin1, bin2)
			totalHd += hd
		}
	}

	return totalHd
}

func IntToBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func padLeft(str string, length int) string {
	if len(str) >= length {
		return str
	}
	return strings.Repeat("0", length-len(str)) + str
}

func HammingDistance(a, b string) int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}

	aPadded := padLeft(a, maxLen)
	bPadded := padLeft(b, maxLen)

	distance := 0
	for i := 0; i < maxLen; i++ {
		if aPadded[i] != bPadded[i] {
			distance++
		}
	}
	return distance
}
