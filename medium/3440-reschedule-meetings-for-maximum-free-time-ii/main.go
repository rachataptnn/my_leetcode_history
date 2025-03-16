package main

import "fmt"

func main() {
	// case 1
	// eventTime := 5
	// startTime := []int{1, 3}
	// endTime := []int{2, 5}

	// case 3
	eventTime := 10
	startTime := []int{0, 3, 7, 9}
	endTime := []int{1, 4, 8, 10}

	fmt.Println(maxFreeTime(eventTime, startTime, endTime))
}

type solver struct {
	startTime  []int
	endTime    []int
	eventTime  int
	boxDetails []boxDetail
	holes      []hole
}

type boxDetail struct {
	longestWithoutThatBox int
	boxLength             int
}

type hole struct {
	start int
	end   int
}

func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	s := solver{
		startTime: startTime,
		endTime:   endTime,
		eventTime: eventTime,
	}

	s.findHoles()
	max := s.calcLongestWithoutThatBox()

	return max
}

func (s *solver) findHoles() {
	// if s.startTime[0] > 0 {
	// 	tmpHole := hole{
	// 		start: 0,
	// 		end:   s.startTime[0],
	// 	}
	// 	s.holes = append(s.holes, tmpHole)
	// }

	for i := 0; i < len(s.startTime)-1; i++ {
		tmpHole := hole{
			start: s.endTime[i],
			end:   s.startTime[i+1],
		}
		s.holes = append(s.holes, tmpHole)
	}
}

func (s *solver) calcLongestWithoutThatBox() int {
	max := 0

	for i := 0; i < len(s.startTime); i++ {
		tmpBox := boxDetail{
			boxLength: s.endTime[i] - s.startTime[i],
		}

		if i == 0 { // first box
			if len(s.startTime) > 1 {
				tmpBox.longestWithoutThatBox = s.startTime[i+1]
			} else {
				tmpBox.longestWithoutThatBox = s.eventTime - s.endTime[i] + tmpBox.boxLength
			}
		} else if i == len(s.startTime)-1 { // last box
			tmpBox.longestWithoutThatBox = s.startTime[i] - s.endTime[i-1] + tmpBox.boxLength
		} else { // mid box
			tmpBox.longestWithoutThatBox = (s.startTime[i+1] - s.endTime[i-1])
		}

		if s.canItJumpToOtherHole(i) {
			tmpBox.longestWithoutThatBox = tmpBox.longestWithoutThatBox
		} else {
			boxLen := s.endTime[i] - s.startTime[i]
			tmpBox.longestWithoutThatBox = tmpBox.longestWithoutThatBox - boxLen
		}

		if tmpBox.longestWithoutThatBox > max {
			max = tmpBox.longestWithoutThatBox
		}

		s.boxDetails = append(s.boxDetails, tmpBox)
	}

	return max
}

func (s *solver) canItJumpToOtherHole(boxNumber int) bool {
	boxStart := s.startTime[boxNumber]
	boxEnd := s.endTime[boxNumber]
	boxLen := boxEnd - boxStart

	for i, hole := range s.holes {
		cantUseThisHole := boxNumber-1 == i || boxNumber == i
		if cantUseThisHole {
			continue
		}

		holeLen := hole.end - hole.start
		canUseThisHole := holeLen >= boxLen
		if canUseThisHole {
			return true
		}
	}

	return false
}
