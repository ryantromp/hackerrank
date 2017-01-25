package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		for j := i; j < n-1; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Printf("#")
		}
		fmt.Println()
	}
}
