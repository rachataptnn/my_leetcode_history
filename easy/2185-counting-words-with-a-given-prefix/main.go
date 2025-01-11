package main

import "fmt"

func main() {
	fmt.Println(prefixCount(nil, ""))
}

func prefixCount(words []string, pref string) int {
	nPrefix := len(pref)
	cnt := 0

	for _, v := range words {
		if len(v) >= nPrefix {
			vPref := v[:nPrefix]
			if vPref == pref {
				cnt++
			}
		}
	}

	return cnt
}
