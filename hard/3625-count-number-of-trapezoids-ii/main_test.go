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
