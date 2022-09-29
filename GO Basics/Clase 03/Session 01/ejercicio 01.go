package main

import (
	"fmt"
	"os"
)


func main(){
	fmt.Println("")
	text := []byte("ID;Precio;Cantidad\n111223;50000;4\n224411;45000;2\n332244;2500;20\n")
	os.WriteFile("ejercicio 01.csv", text, 0644)




	// files, _ := os.ReadDir(".")

	// for _, file := range files {
    //     fmt.Println(file.Name())
    // }
}