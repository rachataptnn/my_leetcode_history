package main

func main() {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"
	// Output: "this is a secret"

	decodeMessage(key, message)
}

func decodeMessage(key string, message string) string {
	m := make(map[string]string)
	alphabets := initAlphabet()

	runningNumber := 0
	for _, v := range key {
		if v == ' ' {
			continue
		}

		_, ok := m[string(v)]
		if !ok {
			m[string(v)] = alphabets[runningNumber]
			runningNumber++
		}

		if runningNumber > 25 {
			break
		}
	}

	output := ""
	for _, v := range message {
		if v == ' ' {
			output += " "
		}

		realLetter := m[string(v)]
		output += string(realLetter)
	}

	return output
}

func initAlphabet() []string {
	var alphabets []string
	for i := 'a'; i <= 'z'; i++ {
		alphabets = append(alphabets, string(i))
	}

	return alphabets
}
