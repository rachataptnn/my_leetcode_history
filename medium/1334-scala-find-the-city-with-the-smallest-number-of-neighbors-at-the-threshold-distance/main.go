// https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/?envType=daily-question&envId=2024-07-26

package main

import (
	"fmt"
)

func main() {
	// n := 4
	// edges := [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}
	// distanceThreshold := 4

	n := 5
	edges := [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}
	distanceThreshold := 2

	// 5/54
	// n := 6
	// edges := [][]int{{0, 3, 7}, {2, 4, 1}, {0, 1, 5}, {2, 3, 10}, {1, 3, 6}, {1, 2, 1}}
	// distanceThreshold := 417

	fmt.Println(findTheCity(n, edges, distanceThreshold))
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	summaryDistances := floydWarshallAlgorithm(edges, n)
	city := getTheRightCity(summaryDistances, distanceThreshold, n)

	return city
}

func floydWarshallAlgorithm(edges [][]int, n int) [][]int {
	summaryDistances := make([][]int, n)
	for i := range summaryDistances {
		summaryDistances[i] = make([]int, n)
		for j := range summaryDistances[i] {
			summaryDistances[i][j] = 10001
		}
		summaryDistances[i][i] = 0
	}

	for _, edge := range edges {
		srcNode := edge[0]
		dstNode := edge[1]
		distance := edge[2]

		summaryDistances[srcNode][dstNode] = distance
		summaryDistances[dstNode][srcNode] = distance
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if summaryDistances[i][k]+summaryDistances[k][j] < summaryDistances[i][j] {
					summaryDistances[i][j] = summaryDistances[i][k] + summaryDistances[k][j]
				}
			}
		}
	}

	return summaryDistances
}

func getTheRightCity(dist [][]int, distanceThreshold, n int) int {
	minReachableCities := n
	result := -1

	for i := 0; i < n; i++ {
		reachableCities := 0
		for j := 0; j < n; j++ {
			if dist[i][j] <= distanceThreshold {
				reachableCities++
			}
		}
		if reachableCities <= minReachableCities {
			minReachableCities = reachableCities
			result = i
		}
	}

	return result
}
