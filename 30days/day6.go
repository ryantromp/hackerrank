package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := scanner.Text()

	numWords, err := strconv.Atoi(n)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	if numWords < 1 || numWords > 10 {
		fmt.Println("Number of words must be between 1 and 10")
		return
	}

	var words []string

	for i := 0; i < numWords; i++ {
		scanner.Scan()
		word := scanner.Text()

		if len(word) == 0 || len(word) > 10000 {
			fmt.Println("word must be between 1 and 10000 characters in length")
		}

		words = append(words, word)
	}

	for _, w := range words {
		result, err := processWord(w)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Println(result)

	}
}

func processWord(word string) (string, error) {

	var odds, evens string = "", ""

	for i := 0; i < len(word); i++ {
		if i%2 == 0 {
			evens += string(word[i])
		} else {
			odds += string(word[i])
		}
	}

	return evens + " " + odds, nil

}
