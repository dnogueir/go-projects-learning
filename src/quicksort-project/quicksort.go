package main

import (
	"fmt"
	"os"
	"strconv"
)

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	orderedNumbers := make([]int, len(numbers))
	copy(orderedNumbers, numbers)

	pivotIndex := len(orderedNumbers) / 2
	pivot := orderedNumbers[pivotIndex]

	orderedNumbers = append(orderedNumbers[:pivotIndex], orderedNumbers[pivotIndex+1:]...)

	bigger, smaller := slicing(orderedNumbers, pivot)

	return append(append(quicksort(smaller), pivot), quicksort(bigger)...)
}

func slicing(numbers []int, pivot int) (bigger []int, smaller []int) {

	for _, n := range numbers {
		if n <= pivot {
			smaller = append(smaller, n)
		} else {
			bigger = append(bigger, n)
		}
	}

	return bigger, smaller
}

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))

	for i, n := range input {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s is not a valid number", n)
		}
		numbers[i] = number
	}

	fmt.Println(quicksort(numbers))
}
