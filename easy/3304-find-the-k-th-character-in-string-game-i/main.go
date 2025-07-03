package main

import (
	"fmt"
)

func main() {
	k := 8 // expect 'b'
	fmt.Println("resut: ", kthCharacter(k))
}

func kthCharacter(k int) byte {
	current := 'a'
	currentStr := string(current)
	var result byte
	for {
		if len(currentStr) > k {
			result = currentStr[k-1]
			break
		}
		if len(currentStr) < k {
			nextStr := convertEveryAlphabetsToNext(currentStr)
			currentStr += nextStr
		}
	}

	return result
}

func convertEveryAlphabetsToNext(in string) string {
	out := ""
	for _, v := range in {
		next := v + 1
		if next > 122 {
			next = 97
		}
		out += string(next)
	}

	return out
}
