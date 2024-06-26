package main

import (
	"testing"
)

func TestMinKBitFlips(t *testing.T) {
	type testCasesStruct struct {
		caseName string
		input    int
		output   int
	}

	testCases := []testCasesStruct{
		{
			caseName: "case 8",
			input:    8,
			output:   6,
		},
		{
			caseName: "case 110",
			input:    110,
			output:   18,
		},
		{
			caseName: "case 385",
			input:    385,
			output:   23,
		},
		{
			caseName: "case 741",
			input:    741,
			output:   35,
		},
		{
			caseName: "case 989",
			input:    989,
			output:   66,
		},
	}

	for _, tt := range testCases {
		t.Run("", func(t *testing.T) {
			result := minSteps(tt.input)
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
