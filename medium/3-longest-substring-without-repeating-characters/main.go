// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

import "fmt"

func main() {
	// s := "pwwkew"

	// 2xx/987
	// s := "au"

	// 279/987
	s := "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	left := 0
	right := 1

	longest := 0
	max := 0
	for right <= len(s) {
		subStr := s[left:right]
		if isAnyCharDup(subStr) {
			left++
			subStr = s[left:right]
			longest = len(subStr) - 1
			fmt.Println("")
		} else {
			right++
			longest++
			if longest > max {
				max = longest
			}
			fmt.Println("")
		}
	}

	return max
}

func isAnyCharDup(subStr string) bool {
	freqMap := make(map[rune]int)
	for _, v := range subStr {
		freqMap[v]++
		if freqMap[v] > 1 {
			return true
		}
	}

	return false
}
