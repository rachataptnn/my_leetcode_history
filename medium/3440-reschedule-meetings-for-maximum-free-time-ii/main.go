package main

import "fmt"

func main() {
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
}

type boxDetail struct {
	longestWithoutThatBox int
	boxLength             int
}

func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	s := solver{
		startTime: startTime,
		endTime:   endTime,
		eventTime: eventTime,
	}

	s.calcLongestWithoutThatBox()

	return 0
}

func (s *solver) calcLongestWithoutThatBox() {
	for i := 0; i < len(s.startTime); i++ {
		tmpBox := boxDetail{
			boxLength: s.endTime[i] - s.startTime[i],
		}

		if i == 0 { // first box
			if len(s.startTime) > 1 {
				tmpBox.longestWithoutThatBox = s.startTime[i+1] - s.endTime[i] + tmpBox.boxLength
			} else {
				tmpBox.longestWithoutThatBox = s.eventTime - s.endTime[i] + tmpBox.boxLength
			}
		} else if i == len(s.startTime)-1 { // last box
			tmpBox.longestWithoutThatBox = s.startTime[i] - s.endTime[i-1] + tmpBox.boxLength
		} else { // mid box
			tmpBox.longestWithoutThatBox = (s.startTime[i+1] - s.endTime[i-1])
		}

		s.boxDetails = append(s.boxDetails, tmpBox)
	}
}
