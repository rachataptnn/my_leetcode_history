package main

func main() { //  0  1  2  3
	nums := []int{1, 3, 5, 6}
	target := 4 // expect 2

	searchInsert(nums, target)
}

func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v > target || v == target {
			return i
		}
	}

	return len(nums)
}
