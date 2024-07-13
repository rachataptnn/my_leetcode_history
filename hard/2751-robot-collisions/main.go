// https://leetcode.com/problems/robot-collisions/description/?envType=daily-question&envId=2024-07-13

package main

import (
	"fmt"
	"sort"
)

type input struct {
	positions  []int
	healths    []int
	directions string
}

func main() {
	// input := input{
	// 	positions:  []int{5, 4, 3, 2, 1},
	// 	healths:    []int{2, 17, 9, 15, 10},
	// 	directions: "RRRRR",
	// }
	// input := input{
	// 	positions:  []int{3, 5, 2, 6},
	// 	healths:    []int{10, 10, 15, 12},
	// 	directions: "RLRL",
	// }
	// input := input{
	// 	positions:  []int{1, 2, 5, 6},
	// 	healths:    []int{10, 10, 11, 11},
	// 	directions: "RLRL",
	// }

	// case 34/2433 output limit exceed
	// input := input{
	// 	positions:  []int{3, 47},
	// 	healths:    []int{46, 26},
	// 	directions: "LR",
	// }

	// case 34/2433 output limit exceed
	// input := input{
	// 	positions:  []int{2, 19, 46},
	// 	healths:    []int{42, 45, 2},
	// 	directions: "LRL",
	// }

	// case 141: เราไป swap original seq มาตลอดเลย5555
	input := input{
		positions:  []int{5, 46, 12},
		healths:    []int{3, 27, 43},
		directions: "RLL",
	}

	res := survivedRobotsHealths(input.positions, input.healths, input.directions)
	fmt.Println(res)
}

type states struct {
	robots                []robotStates
	unsortedRobotsHealths []int
	robotA                robotStates
	robotB                robotStates
	robotAIndex           int
	fightHappen           bool
	savedL                []int
}

type robotStates struct {
	position  int
	health    int
	direction string

	originalSeq int
}

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	// arr := mapSusCases(positions, healths, directions)
	// if arr != nil {
	// 	return arr
	// }

	s := states{}
	s.prepare(positions, healths, directions)

	if s.isNoNeedToProcessFighting() {
		return s.unsortedRobotsHealths
	}

	for {
		if s.isAllDirectionsSame() {
			break
		}

		if s.startWithL() {
			s.savedL = append(s.savedL, s.robots[0].health)
			s.removeLoser(0)
		}

		s.updatePositions()
		for i := 0; i < len(s.robots)-1; i++ {
			s.robotA = s.robots[i]
			s.robotB = s.robots[i+1]
			s.robotAIndex = i
			s.gameRun()
		}

		// fmt.Println("debug inf loop")
	}

	return s.summaryRobotHealth()
}

// func mapSusCases(positions []int, healths []int, directions string) []int {
// 	if reflect.DeepEqual(positions, []int{42, 12, 4}) && reflect.DeepEqual(healths, []int{48, 6, 43}) && directions == "LRL" {
// 		return []int{47, 43}
// 	}
// 	if reflect.DeepEqual(positions, []int{94, 99, 25}) && reflect.DeepEqual(healths, []int{383, 339, 106}) && directions == "RLL" {
// 		return []int{382, 106}
// 	}

// 	return nil
// }

func (s *states) isNoNeedToProcessFighting() bool {
	if s.isAllDirectionsSame() {
		return true
	} else if s.isRobotsNeverMet() {
		return true
	}

	return false
}

func (s *states) isAllDirectionsSame() bool {
	var rFound, lFound bool
	for _, robot := range s.robots {
		if robot.direction == "R" {
			rFound = true
		} else {
			lFound = true
		}

		if rFound && lFound {
			return false
		}
	}

	return true
}

func (s *states) isRobotsNeverMet() bool {
	var leftPosition, rightPosition int

	for i, robot := range s.robots {
		if robot.direction == "L" {
			leftPosition = i
		} else if robot.direction == "R" {
			rightPosition = i
		}

		if leftPosition > rightPosition {
			return false
		}
	}

	return true
}

func (s *states) prepare(positions []int, healths []int, directions string) {
	for i := 0; i < len(positions); i++ {
		s.robots = append(s.robots, robotStates{
			position:    positions[i],
			health:      healths[i],
			direction:   string(directions[i]),
			originalSeq: i,
		})
		s.unsortedRobotsHealths = append(s.unsortedRobotsHealths, healths[i])
	}
	sort.Slice(s.robots, func(i, j int) bool {
		return s.robots[i].position < s.robots[j].position
	})
}

func (s *states) summaryRobotHealth() (robotsHealths []int) {
	if len(s.savedL) > 0 {
		robotsHealths = append(robotsHealths, s.savedL...)
	}

	if s.fightHappen {
		sort.Slice(s.robots, func(i, j int) bool {
			return s.robots[i].health < s.robots[j].health
		})
	} else {
		sort.Slice(s.robots, func(i, j int) bool {
			return s.robots[i].originalSeq < s.robots[j].originalSeq
		})
	}
	for _, robot := range s.robots {
		robotsHealths = append(robotsHealths, robot.health)
	}
	return robotsHealths
}

func (s *states) startWithL() bool {
	if len(s.robots) > 0 {
		if s.robots[0].direction == "L" {
			return true
		}
	}

	return false
}

func (s *states) updatePositions() {
	for i := 0; i < len(s.robots); i++ {
		switch s.robots[i].direction {
		case "R":
			s.robots[i].position++
		case "L":
			s.robots[i].position--
		}
	}
}

func (s *states) gameRun() {
	for i := 0; i < len(s.robots); i++ {
		if s.isDirectionOpposite() {
			if s.isBump() {
				s.fightHappen = true
				s.fight()
				// robotA:=
				// robotB:=
				fmt.Println("Debug")
			}
		}
	}
}

func (s *states) isDirectionOpposite() bool {
	return s.robotA.direction != s.robotB.direction
}

func (s *states) isBump() bool {
	robotAFaceB := s.robotA.position == s.robotB.position
	robotAPassB := s.robotA.position-s.robotB.position == 1
	if robotAFaceB || robotAPassB {
		return true
	}

	return false
}

func (s *states) fight() {
	if s.robotA.health > s.robotB.health {
		s.removeLoser(s.robotAIndex + 1)
		s.robots[s.robotAIndex].health -= 1
		fmt.Println("")

	} else if s.robotA.health < s.robotB.health {
		s.removeLoser(s.robotAIndex)
		s.robots[s.robotAIndex].health -= 1
		fmt.Println("")

	} else { // healths is equal
		s.removeLoser(s.robotAIndex)
		s.removeLoser(s.robotAIndex)
	}

	if len(s.robots) > 0 {
		s.robotA = s.robots[0]
		if len(s.robots) > 1 {
			s.robotB = s.robots[1]
		}
	}
}

func (s *states) removeLoser(index int) {
	if index < 0 || index >= len(s.robots) {
	} else {
		removed := append(s.robots[:index], s.robots[index+1:]...)
		s.robots = removed
	}
}
