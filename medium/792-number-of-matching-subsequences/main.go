package main

import "fmt"

func main() {
	s := "abcde"
	words := []string{"a", "bb", "acd", "ace"}

	fmt.Println(numMatchingSubseq(s, words))
}

func numMatchingSubseq(s string, words []string) int {
	mainMap := generateFreqMap(s)

	cnt := 0
	for _, w := range words {
		subMap := generateFreqMap(w)

		pass := true
		for k, v := range subMap {
			vMainMap, ok := mainMap[k]
			if !ok || vMainMap < v {
				pass = false
				break
			}
		}
		if pass {
			cnt++
		}
	}

	return cnt
}

func generateFreqMap(in string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range in {
		m[v]++
	}

	return m
}
