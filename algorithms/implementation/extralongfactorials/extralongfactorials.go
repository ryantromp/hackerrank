package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	m := big.NewInt(int64(n))
	for n > 1 {
		m.Mul(m, big.NewInt(int64(n-1)))
		n--
	}

	fmt.Printf("%s\n", m.String())
}
