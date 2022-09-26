package main

import "fmt"

func main() {
	edad := 23
	empleo := false
	tiempoTrabajo := 2.0
	sueldo := 100000

	if edad < 22 {
		fmt.Println("Debe ser mayor o igual a 22")
	} else if !empleo {
		fmt.Println("Debe tener empleo acyualmente")
	} else if tiempoTrabajo < 1 {
		fmt.Println("Debe tener más de un año de antigüedad en su empleo")
	} else {
		if sueldo >= 100000 {
			fmt.Println("Está excento de interés")
			fmt.Println("Es apto para recibir crédito")
		}

	}

}
