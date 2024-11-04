// https://leetcode.com/problems/rotate-string/description/

package main

import "fmt"

func main() {
	s := "abcde"
	goal := "cdeab"

	rotateString(s, goal)
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	dg := goal + goal

	for i, v := range dg {
		if v == rune(s[0]) && i < len(s) {
			dgSlice := dg[i : i+len(s)]
			fmt.Println(dgSlice)
			if dgSlice == s {
				return true
			}
		}
	}

	return false
}
