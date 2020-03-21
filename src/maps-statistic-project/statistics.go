package main

import (
	"fmt"
	"os"
	"strings"
)

func getStatistics(words []string) map[string]int {
	statistics := make(map[string]int)

	for _, w := range words {
		firstLetter := strings.ToUpper(string(w[0]))
		counter, found := statistics[firstLetter]
		if found {
			statistics[firstLetter] = counter + 1
		} else {
			statistics[firstLetter] = 1
		}
	}

	return statistics
}

func printMap(statistics map[string]int) {
	fmt.Println("The first letter statistics are:")
	for firstLetter, counter := range statistics {
		if counter > 1 {
			fmt.Printf("Letter %s appeared %d times \n", firstLetter, counter)
		} else {
			fmt.Printf("Letter %s appeared %d time \n", firstLetter, counter)
		}
	}
}

func main() {
	words := os.Args[1:]
	statistics := getStatistics(words)

	printMap(statistics)

}
