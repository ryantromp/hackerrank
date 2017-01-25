package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	sum := big.NewInt(0)
	for i := 0; i < n; i++ {
		bignum := big.NewInt(0)
		fmt.Scanf("%s", bignum)
		sum.Add(sum, bignum)
	}
	fmt.Println(sum.String())
}
