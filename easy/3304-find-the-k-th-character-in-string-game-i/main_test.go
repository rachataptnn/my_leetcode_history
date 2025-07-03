package main

import (
	"fmt"
	"testing"
)

func TestKthCharacter(t *testing.T) {
	testCases := []struct {
		k    int
		want byte
	}{
		{
			k:    5,
			want: 'b',
		},
		{
			k:    10,
			want: 'c',
		},
		{
			k:    8,
			want: 'c',
		},
	}
	for _, tc := range testCases {
		gotC := kthCharacter(tc.k)
		if gotC != tc.want {
			fmt.Printf(` 
======
FAILED
======
kthCharacter(%d) 
got:		%c 
but want: 	%c

`, tc.k, gotC, tc.want)
		}
	}
}
