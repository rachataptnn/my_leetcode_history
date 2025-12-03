package main

import (
	"fmt"
	"math"
)

func main() {
	// points := [][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}} // ex 1

	points := [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}}

	res := countTrapezoids(points)
	fmt.Println(res)
}

func countTrapezoids(points [][]int) int {
	// rip time comp
	pArr := []Point{}
	for _, v := range points {
		tmpP := Point{}
		tmpP[0] = float64(v[0])
		tmpP[1] = float64(v[1])
		pArr = append(pArr, tmpP)
	}

	lines := getLinesFromPoints(pArr)
	fmt.Println(lines)

	cnt := 0
	for _, v := range lines {
		if len(v) > 1 {
			cnt++
		}
	}

	return cnt
}

type Point [2]float64 // index 0 is x

type Line struct {
	Start  Point
	End    Point
	Length float64
	Angle  float64
}

type LinesGroupedByAngle map[float64][]Line

func getLinesFromPoints(points []Point) (lines LinesGroupedByAngle) {
	// initialize the returned map to avoid nil map assignment
	lines = make(LinesGroupedByAngle)
	knewLines := make(map[[4]float64]struct{})

	for _, p1 := range points {
		for _, p2 := range points {
			// get cache
			if isSamePoint(p1, p2) {
				continue
			}
			if isKnewLine(p1, p2, knewLines) {
				continue
			}

			// construct - lines grouped by angle
			line := Line{}.New(p1, p2)
			existedLine, ok := lines[line.Angle]
			if ok {
				existedLine = append(existedLine, line)
				lines[abs(line.Angle)] = existedLine
			} else {
				lines[abs(line.Angle)] = []Line{line}
			}

			// set cache
			p1p2 := [4]float64{p1[0], p1[1], p2[0], p2[1]}
			p2p1 := [4]float64{p2[0], p2[1], p1[0], p1[1]}
			knewLines[p1p2] = struct{}{}
			knewLines[p2p1] = struct{}{}
		}
	}

	return lines
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

// Line - Cache

func isSamePoint(p1, p2 Point) bool {
	p1x, p1y := p1[0], p1[1]
	p2x, p2y := p2[0], p2[1]

	return p1x == p2x && p1y == p2y
}

func isKnewLine(p1, p2 Point, knewLines map[[4]float64]struct{}) bool {
	p1p2 := [4]float64{p1[0], p1[1], p2[0], p2[1]}
	if _, ok := knewLines[p1p2]; ok {
		return true
	}

	return false
}

// Line - Construct

func (l Line) New(point1, point2 Point) Line {
	l.defineStartAndEnd(point1, point2)
	l.getLength()
	l.getAngle()

	return l
}

func (l *Line) defineStartAndEnd(point1, point2 Point) {
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

	l.Start = start
	l.End = end
}

func (l *Line) getLength() {
	x2, x1 := l.End[0], l.Start[0]
	y2, y1 := l.End[1], l.Start[1]

	a := math.Pow(x2-x1, 2)
	b := math.Pow(y2-y1, 2)
	c := math.Sqrt(a + b)

	l.Length = c
}

func (l *Line) getAngle() {
	Ax, Ay := l.Start[0], l.Start[1]
	Bx, By := l.End[0], l.End[1]

	dy := By - Ay
	dx := Bx - Ax
	m := dy / dx

	radians := math.Atan(m)
	l.Angle = radians * 180 / math.Pi
}
