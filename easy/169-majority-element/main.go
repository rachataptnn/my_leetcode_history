package main

func main() {
	nums := []int{3, 2, 3}
	majorityElement(nums)
}

func majorityElement(nums []int) int {
	freqs := make(map[int]int)
	n := len(nums)
	max := 0
	res := nums[0]
	for _, v := range nums {
		freqs[v]++

		if freqs[v]*2 == n {
			return v
		}

		if freqs[v] > max {
			max = freqs[v]
			res = v
		}
	}

	return res
}
