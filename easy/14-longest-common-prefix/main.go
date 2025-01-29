package main

import "fmt"

func main() {
	// strs := []string{"flower", "flow", "flight"}
	// Output: "fl"

	// 80/126
	// strs := []string{"ab", "a"}

	// 118/126
	strs := []string{"reflower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 1 {
		return strs[0]
	}

	shortestWord := ""
	min := 201
	for _, v := range strs {
		if len(v) < min {
			min = len(v)
			shortestWord = v
		}
	}

	for i := 0; i < len(shortestWord); i++ {
		tmpSW := shortestWord[:len(shortestWord)-i]
		nTmpSW := len(tmpSW)
		commonPrefixCnt := 0

		for j := 1; j < n; j++ {
			subWord := strs[j][:nTmpSW]

			if subWord == tmpSW {
				commonPrefixCnt++
			}
		}

		if commonPrefixCnt+1 == len(strs) {
			return tmpSW
		}
	}

	return ""
}
