package main

import (
	"fmt"
	"strconv"
)

func main() {
	// word := "abcde"
	word := "aaaa aaaa aaaa aabb"
	fmt.Println(compressedString(word))
}

func compressedString(word string) string {
	currentChar := word[0]
	currentStack := 1

	comp := ""

	for i := 1; i < len(word); i++ {
		char := word[i]
		if char == currentChar {
			currentStack++
			if currentStack == 9 {
				repeated := strconv.Itoa(currentStack) + string(currentChar)
				comp += repeated

				currentChar = char
				currentStack = 1
			}
		} else {
			repeated := strconv.Itoa(currentStack) + string(currentChar)
			comp += repeated

			currentChar = char
			currentStack = 1
		}
	}

	return comp
}
