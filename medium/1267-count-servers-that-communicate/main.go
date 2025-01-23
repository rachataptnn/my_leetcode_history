package main

import (
	"fmt"
	"strconv"
)

func main() {
	servers := [][]int{{1, 0}, {1, 1}}
	fmt.Println(countServers(servers))
}

func countServers(grid [][]int) int {
	commuServers := make(map[string]struct{})

	for i, row := range grid {
		for j, server := range row {
			if server == 1 {
				if isRowCommu(grid, i) {
					key := strconv.Itoa(i) + ":" + strconv.Itoa(j)
					commuServers[key] = struct{}{}
				} else if isColCommu(grid, j) {
					key := strconv.Itoa(i) + ":" + strconv.Itoa(j)
					commuServers[key] = struct{}{}
				}
			}
		}
	}

	return len(commuServers)
}

func isRowCommu(grid [][]int, i int) bool {
	serverFound := 0
	for _, v := range grid[i] {
		if v == 1 {
			serverFound++
		}
	}

	return serverFound > 1
}

func isColCommu(grid [][]int, j int) bool {
	serverFound := 0
	for i := 0; i < len(grid); i++ {
		if grid[i][j] == 1 {
			serverFound++
		}
	}

	return serverFound > 1
}
