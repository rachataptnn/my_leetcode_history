// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-ii/?envType=daily-question&envId=2024-08-06

package main

import (
	"fmt"
	"sort"
)

func main() {
	word := "aabbccddeeffgghhiiiiii"
	fmt.Println(minimumPushes(word))
}

func minimumPushes(word string) int {
	cntAlp := countAlphabets(word)
	sorted := sortFreqAlphabet(cntAlp)
	pushesCnt := summaryPushCount(sorted)

	return pushesCnt
}

func countAlphabets(word string) map[rune]int {
	cntMap := make(map[rune]int)
	for _, alphabet := range word {
		cntMap[alphabet]++
	}

	return cntMap
}

type alphabetWithFreq struct {
	alphabet rune
	freq     int
}

func sortFreqAlphabet(alpCount map[rune]int) []alphabetWithFreq {
	var sortedSlice []alphabetWithFreq
	for k, v := range alpCount {
		sortedSlice = append(sortedSlice, alphabetWithFreq{k, v})
	}

	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].freq > sortedSlice[j].freq
	})

	return sortedSlice
}

func summaryPushCount(sortedFreq []alphabetWithFreq) int {
	sum := 0
	mul := 1
	for i, v := range sortedFreq {
		if i%8 == 0 && i != 0 {
			mul++
		}
		sum += v.freq * mul
	}

	return sum
}
