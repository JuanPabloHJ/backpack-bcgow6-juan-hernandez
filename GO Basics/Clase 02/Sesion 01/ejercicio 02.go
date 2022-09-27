package main

import (
	"fmt"
	"errors"
)

func promedioNotas (notas ... float64) (promedio float64, err error){
	for _, nota := range notas{
		if nota < 0 {
			err = errors.New("Hay notas negativas")
			return
		}
		promedio+=nota
	}
	promedio/=float64(len(notas))

	return
}

func main (){
	prom, err := promedioNotas(1,2,3,4)
	fmt.Println(prom, err)
}