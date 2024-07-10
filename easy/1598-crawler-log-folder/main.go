// https://leetcode.com/problems/crawler-log-folder/?envType=daily-question&envId=2024-07-10

package main

import "fmt"

func main() {
	input := []string{"d1/", "d2/", "../", "d21/", "./"}
	fmt.Println(minOperations(input))
}

func minOperations(logs []string) int {
	stepsToRoot := 0
	for i := 0; i < len(logs); i++ {
		if logs[i] == "../" { // back
			if stepsToRoot > 0 {
				stepsToRoot--
			}
		} else if logs[i] != "./" { // drill down
			stepsToRoot++
		}
	}
	return stepsToRoot
}
