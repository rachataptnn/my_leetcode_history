// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

package main

import "fmt"

func main() {
	// prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{1, 2}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	max := 0
	for i := len(prices) - 1; i >= 1; i-- {
		sellPrice := prices[i]
		for j := i; j > 0; j-- {
			buyPrice := prices[j-1]
			diff := sellPrice - buyPrice
			if diff > max {
				max = diff
			}
		}
	}

	return max
}
