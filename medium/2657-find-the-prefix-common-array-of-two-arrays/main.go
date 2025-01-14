package main

import "fmt"

func main() {
	A := []int{1, 3, 2, 4}
	B := []int{3, 1, 2, 4}
	// Output: [0,2,3,4]

	fmt.Println(findThePrefixCommonArray(A, B))
}

func findThePrefixCommonArray(A []int, B []int) []int {
	mapA := make(map[int]struct{})
	mapB := make(map[int]struct{})

	res := []int{}

	for i := range A {
		mapA[A[i]] = struct{}{}
		mapB[B[i]] = struct{}{}

		cnt := 0
		for k := range mapA {
			if _, ok := mapB[k]; ok {
				cnt++
			}
		}
		res = append(res, cnt)
	}

	return res
}
