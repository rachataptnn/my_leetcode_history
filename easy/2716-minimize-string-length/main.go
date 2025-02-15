func minimizedStringLength(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}

	return len(m)
}