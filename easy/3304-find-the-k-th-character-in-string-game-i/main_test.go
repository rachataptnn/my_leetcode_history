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
			k:    1,
			want: 'a',
		},
		{
			k:    2,
			want: 'b',
		},
		{
			k:    3,
			want: 'b',
		},
		{
			k:    4,
			want: 'c',
		},
		{
			k:    5,
			want: 'b',
		},
		{
			k:    6,
			want: 'c',
		},
		{
			k:    10,
			want: 'c',
		},
		{
			k:    8,
			want: 'c',
		},
		{
			k:    11,
			want: 'c',
		},

		{
			k:    14,
			want: 'd',
		},
		{
			k:    15,
			want: 'd',
		},
		{
			k:    16,
			want: 'e',
		},

		{
			k:    500,
			want: 'j',
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
