package main

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	type testCase struct {
		nums           []int
		continueRanges []string
	}

	tests := []testCase{
		{
			nums:           []int{0, 2, 3, 4, 6, 8, 9},
			continueRanges: []string{"0", "2->4", "6", "8->9"},
		},
	}

	for _, v := range tests {
		actual := summaryRanges(v.nums)
		if !reflect.DeepEqual(actual, v.continueRanges) {
			t.Errorf(`
				nums: %v,
				actual: %v
				want: %v`, v.nums, actual, v.continueRanges)
		}
	}
}
