// https://leetcode.com/problems/number-of-senior-citizens/?envType=daily-question&envId=2024-08-01

package main

import (
	"fmt"
	"strconv"
)

func main() {
	details := []string{"9751302862F0693", "3888560693F7262", "5485983835F0649", "2580974299F6042", "9976672161M6561", "0234451011F8013", "4294552179O6482"}
	fmt.Println(countSeniors(details))
}

func countSeniors(details []string) int {
	cnt := 0
	for _, person := range details {
		ageInt, _ := strconv.Atoi(person[11:13])
		if ageInt > 60 {
			cnt++
		}
	}

	return cnt
}
