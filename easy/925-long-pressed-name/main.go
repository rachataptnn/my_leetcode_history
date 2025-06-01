package main

import "fmt"

func main() {
	name := "alex"
	typed := "aaleex"
	fmt.Println(isLongPressedName(name, typed))
}

func isLongPressedName(name string, typed string) bool {
	mapName := make(map[string]int)
	mapTyped := make(map[string]int)

	for _, v := range name {
		mapName[string(v)]++
	}
	for _, v := range typed {
		mapTyped[string(v)]++
	}

	for k, vn := range mapName {
		vt, ok := mapTyped[k]
		if !ok || vt < vn {
			return false
		}
	}

	return true
}
