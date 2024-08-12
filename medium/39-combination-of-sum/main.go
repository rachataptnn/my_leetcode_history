// https://leetcode.com/problems/combination-sum/description/

package main

import "fmt"

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7

	fmt.Println(combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int

	if len(candidates) == 0 {
		return result
	}

	backtrack(&result, make([]int, 0), candidates, 0, target)

	return result
}

func backtrack(result *[][]int, currentCombination, candidates []int, startIndex, target int) {
	if target < 0 || startIndex >= len(candidates) {
		return
	}

	if target == 0 { // with `target-candidates[i]` this line will be to when sum of `currentCombination` == target
		combinationCopy := make([]int, len(currentCombination))
		copy(combinationCopy, currentCombination)
		*result = append(*result, combinationCopy)
	}

	for i := startIndex; i < len(candidates); i++ {
		currentCombination = append(currentCombination, candidates[i])
		backtrack(result, currentCombination, candidates, i, target-candidates[i])
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}
