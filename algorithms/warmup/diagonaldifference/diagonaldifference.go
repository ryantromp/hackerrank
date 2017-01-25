package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var matrix [][]int

	for i := 0; i < n; i++ {
		var row []int
		for j := 0; j < n; j++ {
			var el int
			fmt.Scanf("%d", &el)
			row = append(row, el)
		}
		matrix = append(matrix, row)
	}

	var primary, secondary int
	for i := 0; i < n; i++ {
		primary += matrix[i][i]
		secondary += matrix[i][n-i-1]
	}

	diff := math.Abs(float64(primary - secondary))
	fmt.Printf("%d\n", int64(diff))
}
