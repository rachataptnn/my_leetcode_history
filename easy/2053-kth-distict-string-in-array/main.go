// https://leetcode.com/problems/kth-distinct-string-in-an-array/description/?envType=daily-question&envId=2024-08-05

package main

import "fmt"

func main() {
	arr := []string{"d", "b", "c", "b", "c", "a"}
	k := 2

	fmt.Println(kthDistinct(arr, k))
}

func kthDistinct(arr []string, k int) string {
	feqMap := make(map[string]int)
	for _, v := range arr {
		feqMap[v]++
	}

	for _, v := range arr {
		if feqMap[v] == 1 {
			k--
			if k == 0 {
				return v
			}
		}

		// my morning wanna use this On^2 :D
		// for dis, foundCnt := range disMap {
		//     if foundCnt == 1 && v == dis {
		//         cnt++
		//         if cnt == k {
		//             return v
		//         }
		//     }
		// }
	}

	return ""
}
