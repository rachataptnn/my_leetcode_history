// https://leetcode.com/problems/sliding-puzzle/description/

package main

import "fmt"

func main() {
	board := [][]int{{1, 2, 3}, {4, 0, 5}}
	fmt.Println(slidingPuzzle(board))
}

func slidingPuzzle(board [][]int) int {
	s := states{
		board: board,
	}
	s.getZeroPosition()

	for s.swapCnt < 12 {
		zeroIsBottom := s.xZero != 0 && s.yZero == 1
		if zeroIsBottom {
			s.swapLeft()
		}
		if s.isSolve() {
			return s.swapCnt
		}

		zeroUp := s.xZero == 0 && s.yZero == 1
		if zeroUp {
			s.swapTop()
		}
		if s.isSolve() {
			return s.swapCnt
		}

		zeroIsTopLeft := s.xZero != 2 && s.yZero == 0
		if zeroIsTopLeft {
			s.swapRight()
		}
		if s.isSolve() {
			return s.swapCnt
		}

		zeroAtTopRight := s.xZero == 2 && s.yZero == 0
		if zeroAtTopRight {
			s.swapBottom()
		}
		if s.isSolve() {
			return s.swapCnt
		}
	}

	return -1
}

type states struct {
	board   [][]int
	xZero   int
	yZero   int
	swapCnt int
}

func (s *states) getZeroPosition() {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if s.board[i][j] == 0 {
				s.xZero = i
				s.yZero = j
				return
			}
		}
	}
}

func (s *states) swapTop() {
	top := s.board[0][1]

	s.board[0][1] = 0
	s.board[1][1] = top

	s.xZero, s.yZero = 1, 1

	s.swapCnt++
}

func (s *states) swapBottom() {
	bottom := s.board[2][1]

	s.board[2][1] = 0
	s.board[2][0] = bottom

	s.xZero, s.yZero = 2, 1

	s.swapCnt++
}

func (s *states) swapRight() {
	next := s.xZero + 1

	right := s.board[next][0]

	s.board[next][0] = 0
	s.board[s.xZero][0] = right

	s.xZero, s.yZero = next, 0

	s.swapCnt++
}

func (s *states) swapLeft() {
	next := s.xZero - 1

	left := s.board[next][1]

	s.board[next][1] = 0
	s.board[s.xZero][1] = left

	s.xZero, s.yZero = next, 1

	s.swapCnt++
}

func (s *states) isSolve() bool {
	pos1 := s.board[0][0]
	pos2 := s.board[0][1]
	pos3 := s.board[0][2]

	pos4 := s.board[1][0]
	pos5 := s.board[1][1]
	pos6 := s.board[1][2]

	if pos1 == 1 && pos2 == 2 && pos3 == 3 && pos4 == 4 && pos5 == 5 && pos6 == 0 {
		return true
	}
	return false
}
