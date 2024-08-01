// https://leetcode.com/problems/filling-bookcase-shelves/description/?envType=daily-question&envId=2024-07-31

package main

import "fmt"

func main() {
	// expect 6
	// books := [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}
	// shelveWidth := 4

	// expect 4
	// books := [][]int{{1, 3}, {2, 4}, {3, 2}}
	// shelveWidth := 6

	// expect 15
	books := [][]int{{7, 3}, {8, 7}, {2, 7}, {2, 5}}
	shelveWidth := 10

	fmt.Println(minHeightShelves(books, shelveWidth))
}

// TODO: i think the algo would be
// - get the sum of 2 level height:
//   - case 1 lets the Higher next Book stay at the same level
//   - case 2 lets the Higher Next Book place on the upper level
//
// then sumHCase1 and sumHCase2
// choose the lowerH
// do this sub-solution till end of books
func minHeightShelves(books [][]int, shelfWidth int) int {
	// TODO: these compare method would be make testcases pass like 4-5 cases i guess, but for the 20/20 its would be other ways
	minByH := minByH(books, shelfWidth)
	minByW := minByW(books, shelfWidth)

	if minByH < minByW {
		return minByH
	}
	return minByW
}

func minByH(books [][]int, shelfWidth int) int {
	lastBook := books[len(books)-1]
	widthLeft := shelfWidth - lastBook[0]
	maxH, totalH := lastBook[1], lastBook[1]

	for i := len(books) - 2; i > 0; i-- {
		nextBookW := books[i][0]
		nextBookH := books[i][1]

		// isNextLvl := maxH < nextBookH

		if nextBookH <= maxH && widthLeft-nextBookW >= 0 {
			widthLeft -= nextBookW
		} else {
			maxH = nextBookH
			totalH += nextBookH
		}
	}

	totalH += maxH

	return totalH
}

func minByW(books [][]int, shelfWidth int) int {
	lastBook := books[len(books)-1]
	widthLeft := shelfWidth - lastBook[0]
	maxH := lastBook[1]

	totalH := 0

	for i := len(books) - 2; i >= 0; i-- {
		nextBookW := books[i][0]
		nextBookH := books[i][1]

		isLevelHaveSpace := (widthLeft - nextBookW) >= 0
		if isLevelHaveSpace {
			widthLeft -= nextBookW
			if nextBookH > maxH {
				maxH = nextBookH
			}
		} else {
			totalH += maxH
			widthLeft = shelfWidth - nextBookW
			maxH = nextBookH
		}
	}

	totalH += maxH

	return totalH
}
