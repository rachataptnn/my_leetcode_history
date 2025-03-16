package main

func main() {
	isThree(4)
}

func isThree(n int) bool {
	cnt := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			cnt++
		}

		if cnt > 3 {
			return false
		}
	}

	return cnt == 3
}
