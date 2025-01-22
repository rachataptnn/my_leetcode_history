package main

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	rotate(matrix)
}

func rotate(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	arr := make([][]int, rows)
	for i := range arr {
		arr[i] = make([]int, cols)
	}

	n := len(arr) - 1
	for i, row := range matrix {
		for j, cell := range row {
			if n >= i {
				arr[j][n-i] = cell
			}
		}
	}

	for i, row := range arr {
		for j, cell := range row {
			matrix[i][j] = cell
		}
	}
}
