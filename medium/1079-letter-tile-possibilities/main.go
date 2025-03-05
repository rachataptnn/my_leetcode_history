package main

import "fmt"

func main() {
	tiles := "AAB"
	fmt.Println(numTilePossibilities(tiles))
}

func numTilePossibilities(tiles string) int {
	path := []string{}
	freqMap := make(map[string]int)
	var cnt *int

	backtrack(path, freqMap, cnt)
}

func backtrack(path []string, freqMap map[string]int, cnt *int) {
	// Base case: if the path is valid, we count it (every valid sequence)
	// No more letters to add, we stop and count the result
	if len(path) > 0 {
		*cnt++
	}

	// Try every character
	for letter, count := range freqMap {
		if count > 0 {
			path = append(path, letter) // Make the move: add the letter to the path
			freqMap[letter]--           // Update the frequency map (decrement the count for this letter)

			backtrack(path, freqMap, cnt) // Recurse with updated path and freqMap

			path = path[:len(path)-1] // Undo the move (backtrack): remove the letter from path
			freqMap[letter]++         // Restore the frequency map
		}
	}
}
