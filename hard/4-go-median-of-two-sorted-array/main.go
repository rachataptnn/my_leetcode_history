// https://leetcode.com/problems/median-of-two-sorted-arrays/description/

package main

import "fmt"

func main() {
	// 	Example 1:
	// Input: nums1 = [1,3], nums2 = [2]
	// Output: 2.00000
	// Explanation: merged array = [1,2,3] and median is 2.

	// Example 2:
	// Input: nums1 = [1,2], nums2 = [3,4]
	// Output: 2.50000
	// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

	// 2057 / 2094
	// nums1 := []int{1, 3}
	// nums2 := []int{2, 7}

	// 2084 / 2094
	nums1 := []int{3}
	nums2 := []int{-2, -1}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArr := mergeSortedArr(nums1, nums2)
	if (len(nums1)+len(nums2))%2 == 0 {
		leftMedian := float64(mergedArr[(len(mergedArr)/2)-1])
		rightMedian := float64(mergedArr[len(mergedArr)/2])
		return (leftMedian + rightMedian) / 2
	} else {
		return float64(mergedArr[(len(mergedArr) / 2)])
	}
}

func mergeSortedArr(nums1, nums2 []int) []int {
	i, j := 0, 0
	merged := []int{}

	n1 := len(nums1)
	n2 := len(nums2)

	for len(merged) != n1+n2 {
		if j >= n2 {
			merged = append(merged, nums1[i:]...)
		}
		if i >= n1 {
			merged = append(merged, nums2[j:]...)
		}

		if len(merged) == n1+n2 {
			break
		}

		if nums1[i] == nums2[j] {
			merged = append(merged, nums1[i])
			merged = append(merged, nums2[j])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			merged = append(merged, nums2[j])
			j++
		}
	}

	return merged
}
