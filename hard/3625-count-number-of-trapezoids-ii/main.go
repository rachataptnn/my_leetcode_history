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
	// rip time comp
	pArr := []Point{}
	for _, v := range points {
		tmpP := Point{}
		tmpP[0] = v[0]
		tmpP[1] = v[1]
		pArr = append(pArr, tmpP)
	}

	lines := getLinesFromPoints(pArr)
	fmt.Println(lines)

	return 0
}

type Point [2]int // index 0 is x

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
	knewLines := make(map[Point]Point)

	for _, p1 := range points {
		for _, p2 := range points {
			// get cached
			if isSameLine(p1, p2) || isKnewLine(p1, p2, knewLines) {
				continue
			}

			// construct - lines grouped by angle
			line := Line{}.New(p1, p2)
			existedLine, ok := lines[line.Angle]
			if ok {
				existedLine = append(existedLine, line)
				lines[line.Angle] = existedLine

			} else {
				lines[line.Angle] = []Line{line}
			}

			// set cached
			knewLines[p1] = p2
			knewLines[p2] = p1
		}
	}
	return lines
}

// Line - Cache

func isSameLine(p1, p2 Point) bool {
	p1x, p1y := p1[0], p1[1]
	p2x, p2y := p2[0], p2[1]

	return p1x == p2x && p1y == p2y
}

func isKnewLine(p1, p2 Point, knewLines map[Point]Point) bool {
	p2val, ok := knewLines[p1]
	if ok {
		if p2[0] == p2val[0] && p2[1] == p2val[1] {
			return true
		}
	}

	p1val, ok := knewLines[p2]
	if ok {
		if p1[0] == p1val[0] && p1[1] == p1val[1] {
			return true
		}
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
	dx := float64(l.End[0] - l.Start[0])
	dy := float64(l.End[1] - l.Start[1])
	l.Length = math.Hypot(dx, dy)
}

func (l *Line) getAngle() {
	Ax, Ay := l.Start[0], l.Start[1]
	Bx, By := l.End[0], l.End[1]

	term1 := By - Ay
	term2 := Bx - Ax

	m := float64(term1) / float64(term2)

	angle := math.Atan(m)

	l.Angle = angle
}
