package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var pos, neg, zero int

	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		switch {
		case num > 0:
			pos++
		case num < 0:
			neg++
		default:
			zero++
		}
	}

	fmt.Printf("%f\n", float32(pos)/float32(n))
	fmt.Printf("%f\n", float32(neg)/float32(n))
	fmt.Printf("%f\n", float32(zero)/float32(n))
}
