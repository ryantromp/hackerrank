package main

import (
	"fmt"
	"strconv"
)

func main() {

	var time string
	fmt.Scanf("%s\n", &time)

	//Get the AM / PM indicator
	ampm := time[len(time)-2:]
	hour, _ := strconv.Atoi(time[:2])
	switch {
	case ampm == "PM":
		if hour < 12 {
			hour += 12
		}
	default:
		if hour == 12 {
			hour = 0
		}
	}
	fmt.Printf("%02d%s\n", hour, time[2:len(time)-2])
}
