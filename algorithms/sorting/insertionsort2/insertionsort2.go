package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%d\n", &size)

	var arr []int
	for i := 0; i < size; i++ {
		var num int
		fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}

	for i := 1; i < size; i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
		printArray(arr)
	}
}

func printArray(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
