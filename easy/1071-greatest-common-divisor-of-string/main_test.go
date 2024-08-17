package main

import "testing"

func TestSmallestRepeatingUnit(t *testing.T) {
	testCases := []struct {
		full     string
		half     string
		expected string
	}{
		{
			full:     "ABAB",
			half:     "AB",
			expected: "AB",
		},
		{
			full:     "CODE",
			half:     "CO",
			expected: "CODE",
		},
	}

	for _, tc := range testCases {
		result := findCommon(tc.full, tc.half)
		if result != tc.expected {
			t.Errorf(`
			full:     %s
			half:     %s
			expected: %s
			but got:  %s
			`, tc.full, tc.half, tc.expected, result)
		}
	}
}
