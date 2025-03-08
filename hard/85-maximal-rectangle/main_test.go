package main

import (
	"testing"
)

func TestMaximalRectangle(t *testing.T) {
	tests := []struct {
		caseName string
		matrix   [][]byte
		yo       [][]byte
		expected int
	}{
		{
			caseName: "case 1",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			expected: 6,
		},
		{
			caseName: "case 2",
			matrix: [][]byte{
				{'1'},
			},
			expected: 1,
		},
		{
			caseName: "case 3 - 42/74",
			matrix: [][]byte{
				{'0', '1'},
			},
			expected: 1,
		},
		{
			caseName: "case 4 - 44/74",
			matrix: [][]byte{
				{'1', '0', '1', '1', '1'},
				{'0', '1', '0', '1', '0'},
				{'1', '1', '0', '1', '1'},
				{'1', '1', '0', '1', '1'},
				{'0', '1', '1', '1', '1'},
			},
			expected: 6,
		},
		{
			caseName: "case 5 - 69/74",
			matrix: [][]byte{
				{'0', '1', '1', '0', '0', '1', '0', '1', '0', '1'},
				{'0', '0', '1', '0', '1', '0', '1', '0', '1', '0'},
				{'1', '0', '0', '0', '0', '1', '0', '1', '1', '0'},
				{'0', '1', '1', '1', '1', '1', '1', '0', '1', '0'},
				{'0', '0', '1', '1', '1', '1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0', '1', '1', '1', '1', '0'},
				{'0', '0', '0', '1', '1', '0', '0', '0', '1', '0'},
				{'1', '1', '0', '1', '1', '0', '0', '1', '1', '1'},
				{'0', '1', '0', '1', '1', '0', '1', '0', '1', '1'},
			},
			expected: 10,
		},
	}

	for _, test := range tests {
		result := maximalRectangle(test.matrix)
		if result != test.expected {
			t.Errorf(`
			case name: %s 
			expected:  %d 
			but got:   %d`, test.caseName, test.expected, result)
		}
	}
}
