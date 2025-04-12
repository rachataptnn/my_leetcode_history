package main

import "fmt"

func main() {
	pairs := [][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}
	res := validArrangement(pairs)

	fmt.Println(res)
}

func validArrangement(pairs [][]int) [][]int {
	starts := make(map[int][]int)
	ends := make(map[int][]int)
	for i, v := range pairs {
		starts[v[0]] = pairs[i]
		ends[v[1]] = pairs[i]
	}

	initPair := []int{pairs[0][0], pairs[0][1]}
	for _, v := range pairs {
		_, ok := ends[v[0]]
		if !ok {
			initPair = v
		}
	}

	start := initPair[0]
	end := initPair[1]

	result := [][]int{}
	result = append(result, []int{start, end})

	n := len(pairs) - 1
	for n > 0 {
		nextPair := starts[end]
		end = nextPair[1]
		result = append(result, []int{nextPair[0], nextPair[1]})
		n--
	}

	return result
}
