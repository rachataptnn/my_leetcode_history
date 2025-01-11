package main

func main() {
	romanToInt("")
}

var romanMap = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sum, n := 0, len(s)

	for i := 0; i < n; i++ {
		current := romanMap[rune(s[i])]

		if i == n-1 {
			sum += current
			continue
		}

		next := romanMap[rune(s[i+1])]
		if current < next {
			sum += next - current
			i++
			continue
		}

		sum += current
	}

	return sum
}
