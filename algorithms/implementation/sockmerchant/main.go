package main

import "fmt"

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	colors := make(map[int]int)

	var color int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &color)
		colors[color]++
	}

	var pairs int
	for _, val := range colors {
		pairs += val / 2
	}

	fmt.Println(pairs)
}
