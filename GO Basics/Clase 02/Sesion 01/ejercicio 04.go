package main

import (
	// "errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func getMinimum(num1, num2 float64)(minimum float64){
	minimum = num1
	if num2 > num1 {
		minimum = num2
	}
	return
}

func main(){
	fmt.Println(getMinimum(3,2))
}

// func operation (){}

// minFunc, err := operation(minimum)
// averageFunc, err := operation(average)
// maxFunc, err := operation(maximum)
