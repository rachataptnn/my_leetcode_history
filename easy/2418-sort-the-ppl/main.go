// https://leetcode.com/problems/sort-the-people/description/?envType=daily-question&envId=2024-07-22

package main

import "fmt"

func main() {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}

	fmt.Println("Original names:", names)
	sortPeople(names, heights)
	fmt.Println("Sorted names by h:", names)
}

func sortPeople(names []string, heights []int) []string {
	quicksort(heights, names, 0, len(heights)-1)
	return names
}
func quicksort(height []int, names []string, low int, high int) {
	if low < high {
		pi := partition(height, names, low, high)
		quicksort(height, names, low, pi-1)
		quicksort(height, names, pi+1, high)
	}
}

func partition(height []int, names []string, low, high int) int {
	pivot := height[high]
	i := low - 1

	for j := low; j < high; j++ {
		if height[j] > pivot {
			i++
			height[i], height[j] = height[j], height[i]
			names[i], names[j] = names[j], names[i]
		}
	}
	height[i+1], height[high] = height[high], height[i+1]
	names[i+1], names[high] = names[high], names[i+1]
	return i + 1
}
