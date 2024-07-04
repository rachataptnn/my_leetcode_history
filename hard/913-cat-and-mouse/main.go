// https://leetcode.com/problems/cat-and-mouse/description/

package main

import (
	"fmt"
	"sort"
)

type gameStates struct {
	graph         [][]int
	catPosition   int
	mousePosition int
	winner        int

	skipIndexes        []int
	mousePath          []int
	mousePaths         [][]int
	oneStepToHoleNodes []int

	mouseWasPassNode map[int]int
	catWasPassNode   map[int]int
}

func main() {
	// input := [][]int{{2, 5}, {3}, {0, 4, 5}, {1, 4, 5}, {2, 3}, {0, 2, 3}}
	// input := [][]int{{1, 3}, {0}, {3}, {0, 2}}

	// my first invalid case 11/92
	input := [][]int{{3, 4}, {3, 5}, {3, 6}, {0, 1, 2}, {0, 5, 6}, {1, 4}, {2, 4}}

	result := catMouseGame(input)
	fmt.Println("winner: ", result)
}

func catMouseGame(graph [][]int) int {
	gs := gameStates{
		graph:         graph,
		mousePosition: 1,
		catPosition:   2,
		winner:        -1,
	}

	mwp := make(map[int]int)
	cwp := make(map[int]int)
	mwp[1] = 1
	cwp[2] = 1
	gs.mouseWasPassNode = mwp
	gs.catWasPassNode = cwp

	gs.initMousePaths()
	sort2DArrayByLen(gs.mousePaths)

	fmt.Println("MOUSE: ", gs.mousePosition)
	fmt.Println("CAT: ", gs.catPosition)

	round := 0
	for {
		gs.mouseDecidingMove()
		fmt.Println("MOUSE: ", gs.mousePosition)
		if gs.mousePosition == 0 {
			gs.winner = 1
			break
		}

		gs.catDecidingMove()
		fmt.Println("CAT: ", gs.catPosition)

		if gs.winner != -1 {
			break
		}
		// fmt.Println(round)

		round++
	}

	return gs.winner
}

func (gs *gameStates) mouseDecidingMove() {
	if gs.isHoleReachable() {
		gs.winner = 1
	}

	for _, path := range gs.mousePaths {
		if gs.isPathSafe(path) {
			gs.mousePosition = gs.mouseForwardToThisPath(path)
			gs.mouseWasPassNode[gs.mousePosition] += 1
			if gs.mouseWasPassNode[gs.mousePosition] >= 10 {
				gs.winner = 0
			}
			return
		}
	}

	gs.mousePosition = gs.findSafeNode()

}

func (gs *gameStates) isHoleReachable() bool {
	for _, v := range gs.oneStepToHoleNodes {
		if gs.mousePosition == v {
			return true
		}
	}

	return false
}

func (gs *gameStates) isPathSafe(path []int) bool {
	for _, node := range path {
		if node == gs.catPosition {
			return false
		}
	}

	return true
}

func (gs *gameStates) findSafeNode() int {
	safePath := []int{}

	for _, path := range gs.mousePaths {
		for j, node := range path {
			jumpToCat := gs.catPosition == node
			nextRoundCatched := gs.isCatchable(node)
			if jumpToCat || nextRoundCatched {
				break
			}

			if j+1 == len(path) {
				safePath = path
			}
		}
	}

	if len(safePath) > 0 {
		node := gs.mouseForwardToThisPath(safePath)
		return node
	}

	return 0
}

func (gs *gameStates) mouseForwardToThisPath(path []int) int {
	for i, node := range path {
		if node == gs.mousePosition {
			return path[i-1]
		}
	}

	for i, path := range gs.graph {
		if i == gs.mousePosition {
			for _, node := range path {
				if node != gs.catPosition {
					return node
				}
			}
		}
	}

	return 0
}

// CAT DECIDING
func (gs *gameStates) catDecidingMove() {
	if gs.isCatchable(gs.mousePosition) {
		gs.winner = 2
	}

	gs.catPosition = gs.getBlockNode()
	gs.catWasPassNode[gs.catPosition] += 1
	if gs.catWasPassNode[gs.catPosition] >= 10 {
		gs.winner = 0
	}
}

func (gs *gameStates) isCatchable(mousePosition int) bool {
	for _, possibleNode := range gs.graph[gs.catPosition] {
		if possibleNode == mousePosition {
			return true
		}
	}
	return false
}

// BLOCKING STRATEGY
// How we map set to set
// input(mousePosition=4) return 2
// -> ถ้าหนูเดินผ่าน 2 ไป 0 (2 step) lesser!!
// -> ถ้าหนูเดินผ่าน 5 ไป 0 (3 step)

// input(mousePosition=3) return 5
// -> ถ้าหนูเดินผ่าน 2 ไป 0 (3 step)
// -> ถ้าหนูเดินผ่าน 5 ไป 0 (2 step) lesser!!
func (gs *gameStates) getBlockNode() int {
	possiblePath := gs.getMousePossiblePath()
	sort2DArrayByLen(possiblePath)

	predictMousePosition := gs.mouseForwardToThisPath(possiblePath[0])
	return predictMousePosition
}

func (gs *gameStates) getMousePossiblePath() (possiblePath [][]int) {
	for _, path := range gs.mousePaths {
		for _, node := range path {
			if node == gs.mousePosition {
				possiblePath = append(possiblePath, path)
			}
		}
	}

	return possiblePath
}

// INIT: mouse possible paths
func (gs *gameStates) initMousePaths() {
	for i := 0; i < len(gs.graph); i++ {
		foundMap := make(map[int]bool)
		gs.findNextNodes(0, foundMap)
		if len(gs.mousePath) > 0 {
			gs.mousePaths = append(gs.mousePaths, gs.mousePath)
		}
		gs.mousePath = []int{}
	}
}

func (gs *gameStates) findNextNodes(currentTarget int, foundMap map[int]bool) []int {
	for srcIndex, srcNode := range gs.graph {
		for _, desNode := range srcNode {
			if skipThisIndex(srcIndex, gs.skipIndexes) {
				continue
			}
			_, isFound := foundMap[srcIndex]
			if desNode == currentTarget && !isFound {
				if len(gs.mousePath) == 0 { // initiate case
					gs.mousePath = append(gs.mousePath, 0)
					gs.skipIndexes = append(gs.skipIndexes, srcIndex)
					gs.oneStepToHoleNodes = append(gs.oneStepToHoleNodes, srcIndex)
				}

				gs.mousePath = append(gs.mousePath, srcIndex)
				foundMap[currentTarget] = true
				return gs.findNextNodes(srcIndex, foundMap)
			}
		}
	}

	return gs.mousePath
}

func skipThisIndex(srcIndex int, needSkipIndexs []int) bool {
	for _, v := range needSkipIndexs {
		if srcIndex == v {
			return true
		}
	}

	return false
}

func sort2DArrayByLen(arr [][]int) {
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
}
