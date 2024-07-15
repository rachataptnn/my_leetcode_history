package main

import (
	"testing"
)

func TestAyooo(t *testing.T) {
	type testCasesStruct struct {
		caseName string
		input    string
		expected int
	}

	testCases := []testCasesStruct{
		{
			caseName: "leetcode 1",
			input:    "3+2*2",
			expected: 7,
		},
		{
			caseName: "leetcode 2",
			input:    " 3/2 ",
			expected: 1,
		},
		{
			caseName: "leetcode 3",
			input:    " 3+5 / 2 ",
			expected: 5,
		},
		{
			caseName: "leetcode 4",
			input:    "2*3+4",
			expected: 10,
		},

		{
			caseName: "leetcode 5",
			input:    "14/3*2",
			expected: 8, //...
		},
		{
			caseName: "leetcode 6",
			input:    "12/5*5",
			expected: 10,
		},
		{
			caseName: "leetcode 7",
			input:    "1*2-3/4+5*6-7*8+9/10",
			expected: -24,
		},
		{
			caseName: "leetcode 8",
			input:    "1+2*3+4*5+6*7+8*9+10",
			expected: 151,
		},
		{
			caseName: "leetcode 9",
			input:    "1+2*5/3+6/4*2",
			expected: 6, // leetcode expected 6
		},
		{
			caseName: "leetcode 10",
			input:    "123 -8*5 -57/3 +148 +1*3/2*14*11*2*5/4*3/3/3 +2283",
			expected: 2687, // leetcode expected 2623
		},
		{
			caseName: "leetcode 10",
			input:    "530+ 194/2/2*3/25*2/5*6/5*8 -22/2*2*4 +24*11 +120/6/2/2*13*62",
			//         530             22            -88       264      4030
			expected: 4758, // leetcode expected 4752
		},
	}

	for _, tt := range testCases {
		t.Run("", func(t *testing.T) {
			result := calculate(tt.input)
			if result != tt.expected {
				t.Errorf(`
				bro!! 
				test case: %v
				input:     %v
				expected:  %v
				but got:   %v`, tt.caseName, tt.input, tt.expected, result)
			}
		})
	}
}
