// https://leetcode.com/problems/cat-and-mouse/description/

// solution visualized
// https://docs.google.com/spreadsheets/d/1b9A_IwCgyHmhLKBWiE4h2X3O2uwxvchtJfSPDNLcqKc/edit?gid=0#gid=0

package main

import "fmt"

func main() {
	graph := [][]int{
		{3, 4},
		{3, 5},
		{3, 6},
		{0, 1, 2},
		{0, 5, 6},
		{1, 4},
		{2, 4},
	}

	fmt.Println(catMouseGame(graph))
}

// expected output
// 1-mouse
// 2-cat
// 0-draw If ever a position is repeated
func catMouseGame(graph [][]int) int {
	n := len(graph)
	possiblePathCount := (n * n) - n

	s := states{
		graph:           graph,
		possiblePathCnt: possiblePathCount,
	}

	s.initMatrixZeros(len(graph))
	s.fillStepsIntoDP(graph)

	return 0
}

func (s *states) initMatrixZeros(n int) {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	s.dp = matrix
}

type states struct {
	graph [][]int
	dp    [][]int

	possiblePathCnt   int
	calculatedPathCnt int
}

// DP contains the minimum step btw currentNode and destinationNode
func (s *states) fillStepsIntoDP(graph [][]int) [][]int {

	// this would be recursive for sure

	level2Nodes := [][]int{}
	for i, v := range graph {
		level2Nodes = append(level2Nodes, v)
		for _, node := range v {
			if s.dp[i][node] == 0 {
				s.dp[i][node] = 1
			}
		}
	}

	// level3Nodes := [][]int{}
	// for i, v := range level2Nodes {
	// 	level3Nodes = append(level3Nodes, v)
	// 	for _, node := range v {
	// 		if s.dp[i][node] == 0 {
	// 			s.dp[i][node] = 2
	// 		}
	// 	}
	// }

	// level4Nodes := [][]int{}
	// for i, v := range level3Nodes {
	// 	level4Nodes = append(level4Nodes, v)
	// 	for _, node := range v {
	// 		if s.dp[i][node] == 0 {
	// 			s.dp[i][node] = 3
	// 		}
	// 	}
	// }

	// this would be recursive for sure

	return nil
}
