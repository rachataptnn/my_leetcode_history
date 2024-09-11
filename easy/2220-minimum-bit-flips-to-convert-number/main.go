// https://leetcode.com/problems/minimum-bit-flips-to-convert-number/?envType=daily-question&envId=2024-09-11

package main

import (
	"fmt"
	"strconv"
)

func main() {
	start := 10
	goal := 7

	fmt.Println(minBitFlips(start, goal))
}

func minBitFlips(start int, goal int) int {
	sBin := strconv.FormatInt(int64(start), 2)
	gBin := strconv.FormatInt(int64(goal), 2)

	paddedS := fmt.Sprintf("%030s", sBin)
	paddedG := fmt.Sprintf("%030s", gBin)

	flipCnt := 0
	for i := range paddedS {
		if paddedS[i] != paddedG[i] {
			flipCnt++
		}
	}

	return flipCnt
}
