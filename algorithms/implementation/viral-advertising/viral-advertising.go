package main

import "fmt"

func main() {
	var days int
	fmt.Scanf("%d\n", &days)

	mails := 5
	openedMails := 0

	for i := 0; i < days; i++ {
		mails = mails / 2
		openedMails += mails
		mails = mails * 3
	}

	fmt.Println(openedMails)
}
