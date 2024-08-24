// https://leetcode.com/problems/container-with-most-water/

package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height); j++ {
			if i != j {
				w := abs(i - j)
				h := min(height[i], height[j])
				area := w * h
				if area > max {
					max = area
				}
			}
		}
	}

	return max
}

// func initPillars(height []int) []pillar {
// 	res := []pillar{}
// 	for i, h := range height {
// 		res = append(res, pillar{
// 			height: h,
// 			index:  i,
// 		})
// 	}

// 	return res
// }

// type pillar struct {
// 	height int
// 	index  int
// }

// func maxAreaByPillars(pillars []pillar) int {
// 	for i := 0; i < len(pillars); i++ {
// 		for j:=0;
// 	}

// 	return 0
// }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
