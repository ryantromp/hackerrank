package main

import "fmt"

func main() {
	var n, k, q int
	fmt.Scanf("%d %d %d\n", &n, &k, &q)

	var arr []int
	for i := 0; i < n; i++ {
		var elem int
		fmt.Scanf("%d", &elem)
		arr = append(arr, elem)
	}

	if k != n && k != 0 {
		if k > n {
			k = k % n
		}
	}

	var sol []int
	for i := 0; i < q; i++ {
		var qry int
		fmt.Scanf("%d\n", &qry)
		idx := (qry - k + n) % n
		sol = append(sol, arr[idx])
	}

	for _, val := range sol {
		fmt.Printf("%d\n", val)
	}
}
