package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var arr []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}

}
