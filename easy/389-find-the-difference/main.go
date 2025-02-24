package main

import "fmt"

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}

func findTheDifference(s string, t string) byte {
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	for _, v := range s {
		sMap[v]++
	}

	for _, v := range t {
		tMap[v]++
	}

	for key, valT := range tMap {
		valS := sMap[key]
		if valS != valT {
			return byte(key)
		}
	}

	return 0
}
