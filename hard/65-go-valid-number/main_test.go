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
