package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("How to use: conversor <values> <unit>")
		os.Exit(1)
	}

	originUnit := os.Args[len(os.Args)-1]
	originValues := os.Args[1 : len(os.Args)-1]
	var destinyUnit string

	if originUnit == "celsius" {
		destinyUnit = "fahrenheit"
	} else if originUnit == "kilometers" {
		destinyUnit = "miles"
	} else {
		fmt.Printf("%s is not a known unit", destinyUnit)
		os.Exit(1)
	}

	for i, value := range originValues {
		originValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Printf("The value %s in the %d position is not a valid number!\n", value, i)
			os.Exit(1)
		}

		var destinyValue float64

		if originUnit == "celsius" {
			destinyValue = originValue*1.8 + 32
		} else {
			destinyValue = originValue / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s \n", originValue, originUnit, destinyValue, destinyUnit)
	}

}
