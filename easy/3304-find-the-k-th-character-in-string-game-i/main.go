package main

import (
	"fmt"
)

func main() {
	k := 8 // expect 'b'
	fmt.Println("resut: ", kthCharacter(k))
}

func kthCharacter(k int) byte {
	magicNumber := 0
	switch {
	case k == 1:
		return 'a'
	case k == 2, k == 3:
		return 'b'
	case k == 4:
		return 'c'

	case k <= 8: // 2^3
		magicNumber = 1
	case k <= 16:
		magicNumber = 2
	case k <= 32:
		magicNumber = 3
	case k <= 64:
		magicNumber = 4
	case k <= 128:
		magicNumber = 5
	case k <= 256:
		magicNumber = 6
	case k <= 512: // 2^9
		magicNumber = 7
	}

	bruh := 'a' + magicNumber + countMod(k)
	fmt.Println("bruh: ", string(rune(bruh)))

	return byte(bruh)
}

func countMod(in int) int {
	cnt := 0
	for in > 2 {
		in = in / 2
		cnt++
	}

	return cnt
}
