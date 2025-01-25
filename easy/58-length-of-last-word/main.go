package main

import "strings"

func main() {
	lengthOfLastWord("Hello World")
}

func lengthOfLastWord(s string) int {
	split := strings.Split(s, " ")
	for i := len(split) - 1; i > 0; i-- {
		if split[i] != "" {
			return len(split[i])
		}
	}

	return 0
}
