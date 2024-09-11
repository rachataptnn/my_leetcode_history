// https://leetcode.com/problems/maximum-score-words-formed-by-letters/description/

package main

import "fmt"

func main() {
	words := []string{"dog", "cat", "dad", "good"}
	letters := []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}
	score := []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	fmt.Println(maxScoreWords(words, letters, score))
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	letterMap := initLetterMap(score)

	scores := []wordWithScore{}
	for _, word := range words {
		scores = append(scores, wordWithScore{
			word:  word,
			score: scoreOf(word, letterMap),
		})
	}

	return 0
}

type wordWithScore struct {
	word  string
	score int
}

func initLetterMap(score []int) map[rune]int {
	lettersMap := make(map[rune]int)

	i := 0
	for ch := 'a'; ch <= 'z'; ch++ {
		lettersMap[ch] = score[i]
		i++
	}

	return lettersMap
}

func scoreOf(word string, letterMap map[rune]int) int {
	score := 0
	for _, c := range word {
		score += letterMap[c]
	}

	return score
}
