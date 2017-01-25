package main

import "fmt"

func main() {

	var s string
	fmt.Scanf("%s\n", &s)

	count := 1
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			count++
		}
	}

	fmt.Println(count)

}
