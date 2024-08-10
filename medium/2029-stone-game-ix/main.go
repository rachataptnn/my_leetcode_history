// https://leetcode.com/problems/stone-game-ix/

package main

import (
	"fmt"
)

func main() {
	stones := []int{5, 1, 2, 4, 3}
	// stones := []int{2, 1}
	// stones := []int{2}

	// 62/106
	// stones := []int{2, 3}
	fmt.Println(stoneGameIX(stones))
}

func stoneGameIX(stones []int) bool {
	s := states{
		stones: stones,
		round:  1,
		sum:    0,
	}
	isAliceWin := s.pickNum()
	return isAliceWin
}

type states struct {
	stones     []int
	sum        int
	round      int
	isAliceWin bool
}

func (s states) pickNum() bool {
	selectedNum := 0
	selectedIdx := 0
	for i, v := range s.stones {
		isStone3Dividable := (s.sum+v)%3 == 0 // player play optimal
		if !isStone3Dividable {
			selectedIdx = i
			selectedNum = v

			s.sum += v
			s.round++

			s.stones = removeElementByIndex(s.stones, selectedIdx) // the player may remove any stone from stones

			break
		}
	}

	// end case: out of stone, left stone make sum divisible by 3
	if selectedNum == 0 && selectedIdx == 0 {
		noRemainingStones := len(s.stones) == 0
		if noRemainingStones {
			s.isAliceWin = false // Bob will win automatically if: there are no remaining stones
		} else {
			isBobTurn := s.round%2 == 0
			s.isAliceWin = isBobTurn // The player who removes a stone loses if the sum of the values of all removed stones is divisible by 3
		}
		return s.isAliceWin
	}

	return s.pickNum()
}

func removeElementByIndex(arr []int, index int) []int {
	if index >= 0 && index < len(arr) {
		return append(arr[:index], arr[index+1:]...)
	}
	return arr
}
