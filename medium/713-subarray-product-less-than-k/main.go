package main

func main() {
	nums := []int{10, 5, 2, 6}
	k := 100
	numSubarrayProductLessThanK(nums, k)
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	var left, totalCount int
	product := 1

	for right := 0; right < len(nums); right++ {
		product *= nums[right]
		for left <= right && product >= k {
			product /= nums[left]
			left++
		}

		totalCount += right - left + 1
	}

	return totalCount
}
