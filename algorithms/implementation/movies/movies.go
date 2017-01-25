package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i, j, k int
	fmt.Scanf("%d %d %d\n", &i, &j, &k)

	var dayCount int
	for d := i; d <= j; d++ {
		rd := reverseDay(d)
		diff := absoluteDifference(d, rd)
		if diff%k == 0 {
			dayCount++
		}
	}

	fmt.Println(dayCount)

}

func reverseDay(d int) int {
	if d%10 == 0 {
		d = d / 10
	}

	s := strconv.Itoa(d)
	var rs []rune
	for i := len(s) - 1; i >= 0; i-- {
		rs = append(rs, rune(s[i]))
	}

	rd, _ := strconv.Atoi(string(rs))
	return rd
}

func absoluteDifference(d, rd int) int {
	diff := d - rd
	if diff < 0 {
		diff = diff * -1
	}
	return diff
}
