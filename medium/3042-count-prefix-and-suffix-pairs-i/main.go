package main

func main() {
	countPrefixSuffixPairs(nil)
}

func countPrefixSuffixPairs(words []string) int {
	n := len(words)
	cnt := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			word1 := words[i]
			word2 := words[j]

			if len(word1) <= len(word2) {
				prefix := word2[:len(word1)]
				suffix := word2[len(word2)-len(word1):]

				if word1 == prefix && word1 == suffix {
					cnt++
				}
			}
		}
	}

	return cnt
}
