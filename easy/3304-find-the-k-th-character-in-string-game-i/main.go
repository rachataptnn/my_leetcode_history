package main

import (
	"fmt"
	"math"
)

func main() {
	k := 11 // expect 'b'
	fmt.Println("resut: ", kthCharacter(k))
}

func kthCharacter(k int) byte {
	cnt := 0
	// initK := k
	if k == int(math.Pow(2, 2)) ||
		// initK == int(math.Pow(2,3) )||
		k == int(math.Pow(2, 4)) ||
		k == int(math.Pow(2, 5)) ||
		k == int(math.Pow(2, 6)) ||
		k == int(math.Pow(2, 7)) ||
		k == int(math.Pow(2, 8)) ||
		k == int(math.Pow(2, 9)) {
		cnt += 1
	}

	for k > 2 {
		k = k / 2
		cnt++
	}

	fmt.Printf("\n\nk:<%d> cnt:<%d>", k, cnt)
	bruh := 'a' + cnt

	return byte(bruh)
}
