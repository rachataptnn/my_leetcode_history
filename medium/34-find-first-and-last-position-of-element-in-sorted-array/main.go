package main

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8

	searchRange(nums, target)
}

func searchRange(nums []int, target int) []int {
	firstFound := false
	first, last := -1, -1

	for i, v := range nums {
		if v == target && !firstFound {
			firstFound = true
			first = i
			last = i
		} else if v == target {
			last = i
		}
	}

	return []int{first, last}
}
