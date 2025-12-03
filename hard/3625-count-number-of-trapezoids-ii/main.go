package main

import (
	"fmt"
	"math"
)

func main() {
	points := [][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}}

	res := countTrapezoids(points)
	fmt.Println(res)
}

func countTrapezoids(points [][]int) int {
	return 0
}

type Point [2]int // index 0 is x

type Line struct {
	Start  Point
	End    Point
	Length int
	Angle  float64
}

func getLinesFromPoints(points []Pointer) (lines []Line) {
	countedLines := make(map[Point]Point)

	for _, p1 := range points {
		for _, p2 := range points {
			p1x, p1y := p1[0], p1[1]
			p2x, p2y := p2[0], p2[1]
			if p1x == p2x && p1y == p2y {
				continue
			}

			// check if line appended
			point1 := [2]int{p1x, p1y}
			point2 := [2]int{p2x, p2y}
			if isCountedLine(point1, point2) {
				continue
			}

			// define start by compare point
			// define length by some formular lol

			lines = append(lines, Line{}.New(point1, point2))

			// append linep
		}

	}
	return nil
}

func isCountedLine(p1, p2 Point) bool {
	return false
}

func (l Line) New(point1, point2 Point) Line {
	start := Point{}
	end := Point{}

	x1, x2 := point1[0], point2[0]
	isP1BelowP2 := x1 < x2
	if isP1BelowP2 {
		start[0] = point1[0]
		start[1] = point1[1]

		end[0] = point2[0]
		end[1] = point2[1]
	} else {
		end[0] = point1[0]
		end[1] = point1[1]

		start[0] = point2[0]
		start[1] = point2[1]
	}

	somev := math.Pow(end[1] - end[0])
	somv := start[1] - start[0]
	math.Sqrt()

	tmpLine := Line{
		Start:  start,
		End:    end,
		Length: 0,
	}
	tmpLine.getAngle()

	return tmpLine
}

func (l Line) getAngle() {
	Ax, Ay := l.Start[0], l.Start[1]
	Bx, By := l.End[0], l.End[1]

	term1 := By - Ay
	term2 := Bx - Ax

	m := float64(term1) / float64(term2)
	l.Angle = math.Atan(m)
}
