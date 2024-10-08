// https://leetcode.com/problems/cat-and-mouse/description/

// solution visualized
// https://docs.google.com/spreadsheets/d/1b9A_IwCgyHmhLKBWiE4h2X3O2uwxvchtJfSPDNLcqKc/edit?gid=0#gid=0

package main

import (
	"fmt"
	"sort"
)

// Example usage:
func main() {
	graph := [][]int{{2, 5}, {3}, {0, 4, 5}, {1, 4, 5}, {2, 3}, {0, 2, 3}}

	// graph := [][]int{{1, 3}, {0}, {3}, {0, 2}}

	// 11/92
	// graph := [][]int{{3, 4}, {3, 5}, {3, 6}, {0, 1, 2}, {0, 5, 6}, {1, 4}, {2, 4}}

	// 51/92
	// graph := [][]int{{2, 3}, {2}, {0, 1}, {0, 4}, {3}}

	// 33/92
	// graph := [][]int{{2, 3}, {3, 4}, {0, 4}, {0, 1}, {1, 2}}

	fmt.Println(catMouseGame(graph))
}

func buildShortestPathDP(graph [][]int, n int) [][]int {
	// Initialize DP table with maximum values
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if i == j {
				dp[i][j] = 0 // Distance to self is 0
			} else {
				dp[i][j] = n // Initialize with a large value (n is sufficient for unweighted graph)
			}
		}
	}

	// Initialize direct connections (edges) with distance 1
	for node, neighbors := range graph {
		for _, neighbor := range neighbors {
			dp[node][neighbor] = 1
			dp[neighbor][node] = 1 // Since the graph is undirected
		}
	}

	// Floyd-Warshall algorithm
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// If path through k is shorter, update the distance
				if dp[i][k]+dp[k][j] < dp[i][j] {
					dp[i][j] = dp[i][k] + dp[k][j]
				}
			}
		}
	}

	return dp
}

// expected output
// 1-mouse
// 2-cat
// 0-draw If ever a position is repeated
func catMouseGame(graph [][]int) int {
	n := len(graph)
	dp := buildShortestPathDP(graph, n)

	tmpMap := make(map[int]int)
	s := states{
		dp:          dp,
		winner:      -1,
		ratPosition: 1,
		catPosition: 2,

		catFreqVisit: tmpMap,
		ratFreqVisit: tmpMap,
	}
	s.updateDangerNodes()
	s.sortNearZeroNodes()

	for s.winner == -1 {
		s.ratMove()
		if s.winner != -1 {
			break
		}

		s.catMove()
		if s.winner != -1 {
			break
		}

		s.updateDangerNodes()
	}

	return s.winner
}

type states struct {
	dp            [][]int
	nearZeroNodes []vertices
	nearRatNodes  []vertices

	winner int

	ratPosition int
	catPosition int

	dangerNodes map[int]bool

	catFreqVisit map[int]int
	ratFreqVisit map[int]int
}

type vertices struct {
	node  int
	steps int // far from node
}

func (s *states) sortNearZeroNodes() {
	for i, v := range s.dp[0] {
		s.nearZeroNodes = append(s.nearZeroNodes, vertices{
			node:  i,
			steps: v,
		})
	}

	sort.Slice(s.nearZeroNodes, func(i, j int) bool {
		return s.nearZeroNodes[i].steps < s.nearZeroNodes[j].steps
	})
}

func (s *states) updateDangerNodes() {
	dangerNode := map[int]bool{
		s.catPosition: true,
	}
	for node, steps := range s.dp[s.catPosition] {
		if steps == 1 && node != 0 {
			dangerNode[node] = true
		}
	}

	s.dangerNodes = dangerNode
}

// rat func

func (s *states) ratMove() {
	node := s.ratFindNode()
	if node == 0 {
		s.winner = 1
		return
	} else if node == -1 {
		s.winner = 2
		return
	} else {
		s.ratPosition = node
		s.ratFreqVisit[node]++
		if s.catFreqVisit[node] > 3 {
			s.winner = 0
			return
		}
	}
}

func test()

func (s *states) ratFindNode() int {
	possibleNodes := map[int]bool{}
	for node, steps := range s.dp[s.ratPosition] {
		if steps == 1 {
			possibleNodes[node] = true
		}
	}

	for _, node := range s.nearZeroNodes {
		if possibleNodes[node.node] && !s.dangerNodes[node.node] {
			return node.node
		}
	}

	return -1 // cat win?
}

// cat func

func (s *states) catMove() {
	predictedRatNode := s.ratFindNode()
	if predictedRatNode == -1 {
		s.winner = 2
		return
	}

	possibleNodes := map[int]bool{}
	for node, steps := range s.dp[s.catPosition] {
		if steps == 1 && node != 0 {
			possibleNodes[node] = true
		}
	}

	s.nearRatNodes = []vertices{}
	s.sortNearRatNodes(predictedRatNode)

	candidateNodes := make(map[int]vertices)
	firstCandidate := 0
	for _, node := range s.nearRatNodes {
		if possibleNodes[node.node] {
			if len(candidateNodes) > 0 {
				if candidateNodes[firstCandidate].steps == node.steps {
					candidateNodes[node.node] = node
				}
			} else {
				firstCandidate = node.node
				candidateNodes[node.node] = node
			}
		}
	}

	if len(candidateNodes) == 1 {
		s.catPosition = candidateNodes[firstCandidate].node
		s.catFreqVisit[candidateNodes[firstCandidate].node]++
		if s.catFreqVisit[candidateNodes[firstCandidate].node] > 3 {
			s.winner = 0
			return
		}
		return
	} else {
		for _, v := range s.nearZeroNodes {
			_, ok := candidateNodes[v.node]
			if ok {
				s.catPosition = candidateNodes[v.node].node
				s.catFreqVisit[candidateNodes[v.node].node]++
				if s.catFreqVisit[candidateNodes[v.node].node] > 3 {
					s.winner = 0
					return
				}
				return
			}
		}
	}
}

func (s *states) sortNearRatNodes(predictedRatNode int) {
	for i, v := range s.dp[predictedRatNode] {
		s.nearRatNodes = append(s.nearRatNodes, vertices{
			node:  i,
			steps: v,
		})
	}

	sort.Slice(s.nearRatNodes, func(i, j int) bool {
		return s.nearRatNodes[i].steps < s.nearRatNodes[j].steps
	})
}
