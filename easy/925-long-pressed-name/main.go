package main

import "fmt"

func main() {
	// 1 ex true
	// name := "alex"
	// typed := "aaleex"

	// 2 ex false
	name := "saeed"
	typed := "ssaaedd"

	// 76/95 ex false
	// name := "rick"
	// typed := "kric"
	fmt.Println(isLongPressedName(name, typed))
}

func isLongPressedName(name string, typed string) bool {
	n := len(typed)

	index := 0
	for _, nameCh := range name {
		foundNameMatchTyped := false
		for index < n {
			typedCh := rune(typed[index])

			if !foundNameMatchTyped {
				if nameCh == typedCh {
					foundNameMatchTyped = true

					index++
					continue
				} else {
					return false
				}
			} else if nameCh == typedCh {
				index++
				continue
			} else {
				break
				// foundNameMatchTyped = false
				// index++
				// continue
			}
		}
	}

	return true
}

// mapName := make(map[string]int)
// mapTyped := make(map[string]int)

// // isAmountValid
// for _, v := range name {
// 	mapName[string(v)]++
// }
// for _, v := range typed {
// 	mapTyped[string(v)]++
// }
// for k, vn := range mapName {
// 	vt, ok := mapTyped[k]
// 	if !ok || vt < vn {
// 		return false
// 	}
// }

// // isSequenceValid

// return true
