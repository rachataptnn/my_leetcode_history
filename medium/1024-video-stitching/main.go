// https://leetcode.com/problems/video-stitching/description/?envType=problem-list-v2&envId=greedy&difficulty=MEDIUM

package main

import "fmt"

func main() {
	// clips := [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	// time := 10

	// clips := [][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}
	// time := 9

	// clips := [][]int{{0, 4}, {2, 8}}
	// time := 5

	// clips := [][]int{{0, 2}, {4, 8}}
	// time := 5

	clips := [][]int{{0, 5}, {1, 6}, {2, 7}, {3, 8}, {4, 9}, {5, 10}, {6, 11}, {7, 12}, {8, 13}, {9, 14}, {10, 15}, {11, 16}, {12, 17}, {13, 18}, {14, 19}, {15, 20}, {16, 21}, {17, 22}, {18, 23}, {19, 24}, {20, 25}, {21, 26}, {22, 27}, {23, 28}, {24, 29}, {25, 30}, {26, 31}, {27, 32}, {28, 33}, {29, 34}, {30, 35}, {31, 36}, {32, 37}, {33, 38}, {34, 39}, {35, 40}, {36, 41}, {37, 42}, {38, 43}, {39, 44}, {40, 45}, {41, 46}, {42, 47}, {43, 48}, {44, 49}, {45, 50}, {46, 51}, {47, 52}, {48, 53}, {49, 54}}
	time := 50

	fmt.Println(videoStitching(clips, time))
}

func videoStitching(clips [][]int, time int) int {
	s := states{
		currentClip: clipDetail{
			start: clips[0][0],
			end:   clips[0][1],
		},
		n: time,
	}

	for _, clip := range clips {
		s.currentClip.start = clip[0]
		s.currentClip.end = clip[1]

		if s.currentClip.end > s.n {
			s.currentClip.end = s.n
		}
		s.currentClip.length = s.currentClip.end - s.currentClip.start

		if s.isClipBestStart() {
			s.bestStart = s.currentClip
		}

		if s.isClipBestEnd() {
			s.bestEnd = s.currentClip
		}

		if s.isPath() {
			s.paths = append(s.paths, s.currentClip)
		}
	}

	// case -1
	cantReachStartOrEnd := s.bestStart.end == 0 || s.bestEnd.end == 0
	if cantReachStartOrEnd {
		return -1
	}

	// case 1
	startCanReachEnd := s.bestStart.end == time
	endCanReachStart := s.bestEnd.start == 0 && s.bestEnd.end == time
	if startCanReachEnd || endCanReachStart {
		return 1
	}

	// case 2
	bestStartTouchBestEnd := s.bestStart.end >= s.bestEnd.start
	if bestStartTouchBestEnd {
		return 2
	}

	// case n
	pathAmount := s.findMinimumNForConnect()
	if pathAmount != 0 {
		return 2 + pathAmount
	}

	return -1
}

type states struct {
	bestStart clipDetail
	paths     []clipDetail
	bestEnd   clipDetail

	currentClip clipDetail

	n int
}

type clipDetail struct {
	start  int
	end    int
	length int
}

func (s *states) isClipBestStart() bool {
	return s.currentClip.start == 0 && s.currentClip.length > s.bestStart.length
}

func (s *states) isClipBestEnd() bool {
	return s.currentClip.end == s.n && s.currentClip.length > s.bestEnd.length
}

func (s *states) isPath() bool {
	startOrEnd := s.currentClip.start == 0 || s.currentClip.end == s.n
	if startOrEnd {
		return false
	}

	afterBS := s.currentClip.end > s.bestStart.end
	beforeBE := s.currentClip.start < s.bestEnd.start
	if afterBS || beforeBE {
		return true
	}

	return false
}

// case n

func (s *states) findMinimumNForConnect() int {
	candidatePaths := []clipDetail{}

	for _, path := range s.paths {
		if path.start <= s.bestStart.end && path.end >= s.bestEnd.start {
			return 1
		}

		if path.end > s.bestStart.end || path.start < s.bestEnd.start {
			candidatePaths = append(candidatePaths, path)
		}
	}

	return 0
}

// maybe i should take a rest
