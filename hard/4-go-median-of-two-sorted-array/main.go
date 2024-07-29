package main

func main() {
	// 	Example 1:
	// Input: nums1 = [1,3], nums2 = [2]
	// Output: 2.00000
	// Explanation: merged array = [1,2,3] and median is 2.

	// Example 2:
	// Input: nums1 = [1,2], nums2 = [3,4]
	// Output: 2.50000
	// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// n := float64(len(nums1) + len(nums2))
	// if n%2 == 0 {

	// 	return 0
	// }

	// 2057 / 2094 testcases passed with these line below lol, anyway median is not equal with average :(
	sum := 0
	for _, v := range nums1 {
		sum += v
	}
	for _, v := range nums2 {
		sum += v
	}

	n := float64(len(nums1) + len(nums2))
	res := float64(sum) / n
	return res
}
