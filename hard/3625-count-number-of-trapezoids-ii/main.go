package main

import (
	"fmt"
	"math"
)

func main() {
	// points := [][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}} // ex 1
	// points := [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}} // ex 2

	// points := [][]int{{-32, 12}, {-32, -94}, {-32, -15}, {-30, 88}} //
	points := [][]int{{71, -89}, {-75, -89}, {-9, 11}, {-24, -89}, {-51, -89}, {-77, -89}, {42, 11}} // 418/551

	// points := [][]int{{-33, -9}, {30, -37}, {-10, -9}, {61, -9}, {56, -67}, {36, -9}, {36, 100}, {36, 96}, {-32, 84}, {18, 34}, {-10, -82}} // 482/551

	res := countTrapezoids(points)
	fmt.Println("\n", res)
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
	// rip time comp

	lines := getLinesFromPoints(pArr)

	cnt := countUniqTraps(lines)

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

	y1, y2 := point1[1], point2[1]
	isP1BelowP2 := y1 < y2
	if isP1BelowP2 {
		start[0], start[1] = point1[0], point1[1]
		end[0], end[1] = point2[0], point2[1]
	} else {
		start[0], start[1] = point2[0], point2[1]
		end[0], end[1] = point1[0], point1[1]
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
	degree := radians * 180 / math.Pi

	if radians < 0 {
		degree = 180 + degree
	}
	l.Angle = degree
}

// find unique trap

func countUniqTraps(linesGroupedByAngle LinesGroupedByAngle) int {
	cnt := 0
	for angle, lines := range linesGroupedByAngle {
		checkedTraps := make(checkedTraps)

		for i, l1 := range lines {
			for j, l2 := range lines {
				if i == j {
					continue
				}
				if checkedTraps.isChecked(l1, l2) { // get cache
					continue
				}

				if isTrapezoid(l1, l2) {
					// fmt.Printf("\nAngle[%2f] - found Trapezoid using l1:<%+v> and l2:<%+v>", angle, l1, l2)
					cnt++
					fmt.Printf("\nFound:%d - angle:%.2f\n", cnt, angle)
					fmt.Println(formatEasyToReadLine(l1, l2))
				}

				checkedTraps.setCheckedTrap(l1, l2) // set cache
			}
		}
	}

	return cnt
}

type checkedTraps map[[8]float64]struct{}

func (c *checkedTraps) isChecked(l1, l2 Line) bool {
	if l1.Angle == 0.00 || l1.Angle > 56 {
		if l1.Start[0] == -24.00 || l2.Start[0] == -24.00 {
			fmt.Println("here")
		}
	}

	checkedTraps := *c

	// l1
	p1x, p1y := l1.Start[0], l1.Start[1]
	p2x, p2y := l1.End[0], l1.End[1]
	// l2
	p3x, p3y := l2.Start[0], l2.Start[1]
	p4x, p4y := l2.End[0], l2.End[1]

	// line swap overcount - prevention
	p1p2p3p4 := [8]float64{
		p1x, p1y, p2x, p2y,
		p3x, p3y, p4x, p4y}
	if _, ok := checkedTraps[p1p2p3p4]; ok {
		return true
	}

	// parallelogram overcounts - prevention
	p1p3p2p4 := [8]float64{
		p1x, p1y, p3x, p3y,
		p2x, p2y, p4x, p4y}
	if _, ok := checkedTraps[p1p3p2p4]; ok {
		return true
	}

	return false
}

func (c *checkedTraps) setCheckedTrap(l1, l2 Line) {
	checkedTraps := *c

	// l1
	p1x, p1y := l1.Start[0], l1.Start[1]
	p2x, p2y := l1.End[0], l1.End[1]
	// l2
	p3x, p3y := l2.Start[0], l2.Start[1]
	p4x, p4y := l2.End[0], l2.End[1]

	// line swap overcount - prevention
	p1p2p3p4 := [8]float64{
		p1x, p1y, p2x, p2y,
		p3x, p3y, p4x, p4y}
	p3p4p1p2 := [8]float64{
		p3x, p3y, p4x, p4y,
		p1x, p1y, p2x, p2y}
	checkedTraps[p1p2p3p4] = struct{}{}
	checkedTraps[p3p4p1p2] = struct{}{}

	// parallelogram overcounts - prevention
	p1p3p2p4 := [8]float64{
		p1x, p1y, p3x, p3y,
		p2x, p2y, p4x, p4y}
	p2p4p1p3 := [8]float64{
		p2x, p2y, p4x, p4y,
		p1x, p1y, p3x, p3y}
	checkedTraps[p1p3p2p4] = struct{}{}
	checkedTraps[p2p4p1p3] = struct{}{}

	c = &checkedTraps
}

func isTrapezoid(l1, l2 Line) bool {
	x1, y1 := l1.Start[0], l1.Start[1]
	x2, y2 := l1.End[0], l1.End[1]
	x3, y3 := l2.Start[0], l2.Start[1]

	term1 := y2 - y3
	term2 := y3 - y1
	term3 := y1 - y2
	areaOf3Dots := 0.5 * (x1*(term1) + x2*(term2) + x3*(term3))

	return abs(areaOf3Dots) > 0
}

// printer helping
func formatEasyToReadLine(l1, l2 Line) string {
	text := `
	line1 [%7.2f, %7.2f] -> [%7.2f, %7.2f]
	line2 [%7.2f, %7.2f] -> [%7.2f, %7.2f]`

	return fmt.Sprintf(
		text,
		l1.Start[0], l1.Start[1], l1.End[0], l1.End[1],
		l2.Start[0], l2.Start[1], l2.End[0], l2.End[1],
	)
}
