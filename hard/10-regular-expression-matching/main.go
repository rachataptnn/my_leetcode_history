// https://leetcode.com/problems/regular-expression-matching/description/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// 354/356
	a := "abc"
	p := "a***abc"

	// 354/356 again
	// a := "a"
	// p := ".*..a*"

	fmt.Println(isMatch(a, p))
}

// ppl doing dp, dfs, recursion
// feel lil bit guilty to solve this prob using lib ._.

func isMatch(s string, p string) bool {
	normalizedPattern := normalizePattern(p)

	// Add anchors to ensure the pattern matches the entire string
	pattern := "^" + normalizedPattern + "$"

	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}

func normalizePattern(p string) string {
	var sb strings.Builder
	prevChar := byte(0)

	for i := 0; i < len(p); i++ {
		if p[i] == '*' && prevChar == '*' {
			continue // Skip consecutive '*'
		}
		sb.WriteByte(p[i])
		prevChar = p[i]
	}

	return sb.String()
}
