// https://leetcode.com/problems/pass-the-pillow/?envType=daily-question&envId=2024-07-16

package main

import "fmt"

func main() {
	// n, time := 4, 5
	// n, time := 8, 9

	n, time := 18, 38
	fmt.Println(passThePillow(n, time))

}

type states struct {
	stepsLeft        int
	currentPosition  int
	endpoint         int
	countInThatRound int
}

func passThePillow(n int, time int) int {
	s := states{
		stepsLeft:       time,
		currentPosition: 1,
		endpoint:        n,
	}

	for s.stepsLeft != 0 {
		s.forward()
		if s.stepsLeft == 0 {
			s.countInThatRound += 1
			break
		}
		s.backward()
	}

	return s.countInThatRound
}

func (s *states) forward() {
	s.countInThatRound = 0
	for {
		if s.currentPosition == s.endpoint || s.stepsLeft == 0 {
			break
		}

		s.stepsLeft--
		s.countInThatRound++
		s.currentPosition++
	}
}

func (s *states) backward() {
	s.countInThatRound = s.endpoint
	for {
		if s.currentPosition == 1 || s.stepsLeft == 0 {
			break
		}

		s.stepsLeft--
		s.countInThatRound--
		s.currentPosition--
	}
}
