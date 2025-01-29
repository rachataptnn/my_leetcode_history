package main

import "fmt"

func main() {
	// strs := []string{"flower", "flow", "flight"}
	// Output: "fl"

	// 80/126
	strs := []string{"ab", "a"}

	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	firstWord := strs[0]
	n := len(strs)
	if n == 1 {
		return firstWord
	}

	for i := 0; i < len(firstWord); i++ {
		tmpFW := firstWord[:len(firstWord)-i]
		nTmpFW := len(tmpFW)
		commonPrefixCnt := 0

		for j := 1; j < n; j++ {
			// if len(strs[j]) != len(tmpFW) {
			subWord := ""
			if len(strs[j]) > nTmpFW {
				subWord = strs[j][:nTmpFW]
			} else if len(strs[j]) < nTmpFW {
				continue
			}

			if subWord == tmpFW {
				commonPrefixCnt++
			}
		}

		if commonPrefixCnt+1 == len(strs) {
			return tmpFW
		}
	}

	return ""
}
