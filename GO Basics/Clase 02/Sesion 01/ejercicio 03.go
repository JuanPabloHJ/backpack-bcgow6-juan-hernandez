package main

import (
	"fmt"
	"errors"
)

const (
	A = "A"
	B = "B"
	C = "C"
)

func calculateSalary(minutes float64, category string)(finalSalary float64, err error){
	finalSalary=minutes/60
	if minutes < 0 {
		err = errors.New("Numéro de minutos negativo")
	}
	if category == C{
		finalSalary*=1000
	} else if category == B{
		finalSalary*=1500
		finalSalary+=finalSalary*0.2
	} else if category == A{
		finalSalary*=3000
		finalSalary+=finalSalary*0.5
	} else {
		err = errors.New("Categoría erronea")
	}

	return
}

func main (){
	fmt.Println(calculateSalary(-600,C))
}