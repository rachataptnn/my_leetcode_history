package main

func main() {
	nums := []int{10, 5, 2, 6}
	k := 100
	numSubarrayProductLessThanK(nums, k)
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	cnt := 0
	n := len(nums)
	for start := 0; start < n; start++ {
		for end := start; end < n; end++ {
			subarray := nums[start : end+1]

			product := 1
			isOver := false
			for _, v := range subarray {
				product *= v
				if product >= k {
					isOver = true
					break
				}
			}
			if !isOver {
				cnt++
			}
		}
	}

	return cnt
}
