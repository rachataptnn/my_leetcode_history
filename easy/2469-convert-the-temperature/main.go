package main

import "fmt"

func main() {
	celsius := 36.50
	res := convertTemperature(celsius)
	fmt.Print(res)
}

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.80 + 32.00}
}
