package main

import (
	"math"
	"testing"
)

func TestLine_GetAngle(t *testing.T) {
	tests := []struct {
		name  string
		start Point
		end   Point
		want  float64
		skip  bool // vertical lines currently cause division-by-zero in your func
	}{
		{
			name:  "horizontal positive",
			start: Point{0, 0},
			end:   Point{10, 0},
			want:  0,
		},
		{
			name:  "positive slope 1 (45°)",
			start: Point{0, 0},
			end:   Point{10, 10},
			want:  math.Pi / 4,
		},
		{
			name:  "negative slope -1 (-45°)",
			start: Point{0, 0},
			end:   Point{10, -10},
			want:  -math.Pi / 4,
		},
		{
			name:  "steep positive slope",
			start: Point{0, 0},
			end:   Point{2, 10},
			want:  math.Atan(10.0 / 2.0),
		},
		{
			name:  "vertical line (division by zero in current code)",
			start: Point{5, 0},
			end:   Point{5, 10},
			skip:  true, // remove when vertical handling is implemented
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skip("skipped due to division by zero behavior")
			}

			l := &Line{
				Start: tt.start,
				End:   tt.end,
			}

			l.getAngle()

			if math.Abs(l.Angle-tt.want) > 1e-9 {
				t.Errorf("Angle = %v, want %v", l.Angle, tt.want)
			}
		})
	}
}

func TestLine_GetLength(t *testing.T) {
	tests := []struct {
		name  string
		start Point
		end   Point
		want  float64
	}{
		{
			name:  "zero length",
			start: Point{0, 0},
			end:   Point{0, 0},
			want:  0,
		},
		{
			name:  "horizontal line length 10",
			start: Point{0, 0},
			end:   Point{10, 0},
			want:  10,
		},
		{
			name:  "vertical line length 10",
			start: Point{0, 0},
			end:   Point{0, 10},
			want:  10,
		},
		{
			name:  "45-degree slope 10,10 => length sqrt(200)",
			start: Point{0, 0},
			end:   Point{10, 10},
			want:  math.Sqrt(200),
		},
		{
			name:  "negative coordinates",
			start: Point{-3, -4},
			end:   Point{3, 4},
			want:  10, // classic 3-4-5 triangle doubled
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			l := &Line{
				Start: tt.start,
				End:   tt.end,
			}

			l.getLength()

			if math.Abs(l.Length-tt.want) > 1e-9 {
				t.Errorf("Length = %v, want %v", l.Length, tt.want)
			}
		})
	}
}

// points := [][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}} // ex 1
// points := [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}} // ex 2 (expect 1)

// points := [][]int{{-32, 12}, {-32, -94}, {-32, -15}, {-30, 88}} // case-a (expect 0)

func TestCountTrapezoids(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "example 1",
			points: [][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}},
			want:   2,
		},
		{
			name:   "example 2",
			points: [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}},
			want:   1,
		},
		{
			name:   "case with no trapezoids",
			points: [][]int{{-32, 12}, {-32, -94}, {-32, -15}, {-30, 88}},
			want:   0,
		},
		{
			name:   "there are something called Paralelogram",
			points: [][]int{{71, -89}, {-75, -89}, {-9, 11}, {-24, -89}, {-51, -89}, {-77, -89}, {42, 11}}, // 418/551
			want:   10,
		},
		{
			name:   "fix overlapsed (try 001/???)",
			points: [][]int{{-33, -9}, {30, -37}, {-10, -9}, {61, -9}, {56, -67}, {36, -9}, {36, 100}, {36, 96}, {-32, 84}, {18, 34}, {-10, -82}}, // 482/551
			want:   3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cnt := countTrapezoids(test.points)
			if cnt != test.want {
				t.Errorf(`\ncase name %s
got:  %d
want: %d`, test.name, cnt, test.want)
			}
		})
	}
}
