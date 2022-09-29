package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func getMinimum(numbers ...float64) (minimum float64) {
	minimum = math.Inf(1)

	for _, number := range numbers {
		if number < minimum {
			minimum = number
		}
	}
	return
}

func getMaximum(numbers ...float64) (maximum float64) {
	maximum = math.Inf(-1)
	for _, number := range numbers {
		if number > maximum {
			maximum = number
		}
	}
	return
}

func getAverage(numbers ...float64) (average float64) {
	average = 0
	for _, number := range numbers {
		average += number
	}
	return average / float64(len(numbers))
}

func operation(operationName string) (function func(...float64) float64, err error) {
	switch operationName {
	case minimum:
		function = getMinimum
	case maximum:
		function = getMaximum
	case average:
		function = getAverage
	default:
		err = errors.New("unknown operation: " + operationName)
	}
	return
}

func main() {
	maxFunction, _ := operation(maximum)
	minFunction, _ := operation(minimum)
	avgFunction, _ := operation(average)

	fmt.Println(maxFunction(1, 2, 3, 4, 5))
	fmt.Println(minFunction(1, 2, 3, 4, 5))
	fmt.Println(avgFunction(1, 2, 3, 4, 5))
}
