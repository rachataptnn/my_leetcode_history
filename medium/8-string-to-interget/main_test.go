package main

import (
	"testing"
)

func TestMyAtoiFlips(t *testing.T) {
	type testCasesStruct struct {
		caseName string
		input    string
		output   int
	}

	testCases := []testCasesStruct{
		{
			caseName: "default case from Leetcode 1",
			input:    "42",
			output:   42,
		},
		{
			caseName: "default case from Leetcode 2",
			input:    " -42",
			output:   -42,
		},
		{
			caseName: "default case from Leetcode 3",
			input:    "1337c0d3",
			output:   1337,
		},
		{
			caseName: "default case from Leetcode 4",
			input:    "0-1",
			output:   0,
		},
		{
			caseName: "default case from Leetcode 5",
			input:    "words and 987",
			output:   0,
		},
		{
			caseName: "edge case 32 bit positive",
			input:    "2147483649",
			output:   2147483647,
		},
		{
			caseName: "edge case 32 bit negative",
			input:    "-2147483649",
			output:   -2147483648,
		},
		{
			caseName: "edge case positive symbolic",
			input:    "+1",
			output:   1,
		},
		{
			caseName: "edge case so longgg positive",
			input:    "20000000000000000000",
			output:   2147483647,
		},
	}

	for _, tt := range testCases {
		t.Run("", func(t *testing.T) {
			result := myAtoi(tt.input)
			if result != tt.output {
				t.Errorf(`
				case name: %s
				input:     %v 
				output:    %v`, tt.caseName, tt.input, tt.output)
			}
		})
	}
}
