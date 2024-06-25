package main

import (
	"testing"
)

func TestMinKBitFlips(t *testing.T) {
	type testCasesStruct struct {
		caseName string
		nums     []int
		k        int
		expected int
	}

	testCases := []testCasesStruct{
		{
			caseName: "default case from Leetcode 1",
			nums:     []int{0, 1, 0},
			k:        1,
			expected: 2,
		},
		{
			caseName: "default case from Leetcode 2",
			nums:     []int{1, 1, 0},
			k:        2,
			expected: -1,
		},
		{
			caseName: "default case from Leetcode 3",
			nums:     []int{0, 0, 0, 1, 0, 1, 1, 0},
			k:        3,
			expected: 3,
		},
	}

	// {
	// 	// Test cases
	// 	{[]int{0, 1, 0}, 1, 2},
	// 	{[]int{1, 1, 0}, 2, -1},
	// 	{[]int{0, 0, 0, 1, 0, 1, 1, 0}, 3, 3},
	// 	{[]int{1, 0, 1, 0, 1, 0, 1}, 3, 2},
	// 	{[]int{0, 0, 0, 0, 0}, 2, 3},
	// 	{[]int{0, 0, 0, 0, 0}, 3, 2},
	// 	{[]int{1, 1, 1, 1, 1}, 1, 0},
	// 	{[]int{1, 0, 0, 0, 1, 0}, 3, -1},
	// }

	for _, tt := range testCases {
		t.Run("", func(t *testing.T) {
			result := minKBitFlips(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("minKBitFlips(%v, %d) = %d; want %d", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
