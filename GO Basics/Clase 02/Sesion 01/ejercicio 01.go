package main

import (
	"fmt"
	"errors"
)

func impuesto(salario float64) (impuesto float64, err error){
	if salario<0{
		err = errors.New("Salario negativo")
	}
	if salario > 50_000 {
		impuesto+=salario*0.17
	}
	if salario > 150_000 {
		impuesto+=salario*0.10
	}

	return

}

func main (){

	fmt.Println(impuesto(-20000))
	fmt.Println(impuesto(20000))
	fmt.Println(impuesto(60000))
	fmt.Println(impuesto(200000))
	

}