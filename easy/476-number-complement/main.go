// https://leetcode.com/problems/number-complement/?envType=daily-question&envId=2024-08-22

package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 5
	fmt.Println(findComplement(num))
}

func findComplement(num int) int {
	binNum := intToBinary(num)
	flipped := flipBits(binNum)
	res := binaryToInt(flipped)

	return res
}

func intToBinary(num int) string {
	binary := ""
	for num > 0 {
		binary = fmt.Sprintf("%d%s", num%2, binary)
		num /= 2
	}
	return binary
}

func flipBits(binary string) string {
	flipped := ""
	for _, bit := range binary {
		if bit == '0' {
			flipped += "1"
		} else {
			flipped += "0"
		}
	}
	return flipped
}

func binaryToInt(binary string) int {
	num, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}
