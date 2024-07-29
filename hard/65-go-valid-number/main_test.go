package main

import (
	"testing"
)

func TestFn(t *testing.T) {
	type testCasesStruct struct {
		caseName string
		input    string
		output   bool
	}

	testCases := []testCasesStruct{
		{
			caseName: "default leetcode - 1",
			input:    "0",
			output:   true,
		},
		{
			caseName: "default leetcode - 2",
			input:    "e",
			output:   false,
		},
		{
			caseName: "default leetcode - 3",
			input:    ".",
			output:   false,
		},

		{
			caseName: "edge case leetcode - 1 (967/1494)",
			input:    ".1",
			output:   true,
		},
		{
			caseName: "edge case leetcode - 2 (1231/1494)",
			input:    "2e0",
			output:   true,
		},
		{
			caseName: "edge case leetcode - 3 (1298/1494)",
			input:    "e9",
			output:   false,
		},
		{
			caseName: "edge case leetcode - 4 (1331/1494)",
			input:    "0e",
			output:   false,
		},
		{
			caseName: "edge case leetcode - 5 (1361/1494)",
			input:    ".e1",
			output:   false,
		},
		{
			caseName: "edge case leetcode - 6 (1355/1494)",
			input:    "1e.",
			output:   false,
		},
		{
			caseName: "edge case leetcode - 7 (1375/1494)",
			input:    "te1",
			output:   false,
		},
		{
			caseName: "edge case leetcode - 8 (1306/1494)",
			input:    "3.",
			output:   true,
		},

		{
			caseName: "edge case leetcode - 9 (1368/1494)",
			input:    ".-4",
			output:   false,
		},
		{
			caseName: "fkcname",
			input:    "+.8",
			output:   true,
		},
		{
			caseName: "fkcname",
			input:    "+.",
			output:   false,
		},
		{
			caseName: "fkcname",
			input:    ".+8",
			output:   false,
		},

		{
			caseName: "fkcname",
			input:    ".0e7",
			output:   true,
		},

		{
			caseName: "fkcname",
			input:    "6ee69",
			output:   false,
		},
		{
			caseName: "fkcname",
			input:    "3E+7",
			output:   false,
		},
	}

	for _, tt := range testCases {
		t.Run("", func(t *testing.T) {
			result := isNumber(tt.input)
			if result != tt.output {
				t.Errorf(`case: %s
					input:    %v
					expected: %v
					but got:  %v
				`, tt.caseName, tt.input, tt.output, result)
			}
		})
	}
}
