package main

func main() {
	nums := []int{1, -3, 2, 3, -4}
	maxAbsoluteSum(nums)
}

func maxAbsoluteSum(nums []int) int {
	curPositiveSum, curNegativeSum := 0, 0
	maxAbsSum := 0

	for _, num := range nums {
		curPositiveSum = max(0, curPositiveSum+num)
		curNegativeSum = max(0, curNegativeSum-num)
		maxAbsSum = max(maxAbsSum, max(curPositiveSum, curNegativeSum))
	}

	return maxAbsSum
}
