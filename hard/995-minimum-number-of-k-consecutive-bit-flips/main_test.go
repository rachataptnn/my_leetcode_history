package main

import (
	"testing"
)

func TestMinKBitFlips(t *testing.T) {
	tests := []struct {
		A        []int
		K        int
		expected int
	}{
		// Test cases
		{[]int{0, 1, 0}, 1, 2},
		{[]int{1, 1, 0}, 2, -1},
		{[]int{0, 0, 0, 1, 0, 1, 1, 0}, 3, 3},
		{[]int{1, 0, 1, 0, 1, 0, 1}, 3, 2},
		{[]int{0, 0, 0, 0, 0}, 2, 3},
		{[]int{0, 0, 0, 0, 0}, 3, 2},
		{[]int{1, 1, 1, 1, 1}, 1, 0},
		{[]int{1, 0, 0, 0, 1, 0}, 3, -1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := minKBitFlips(tt.A, tt.K)
			if result != tt.expected {
				t.Errorf("minKBitFlips(%v, %d) = %d; want %d", tt.A, tt.K, result, tt.expected)
			}
		})
	}
}
