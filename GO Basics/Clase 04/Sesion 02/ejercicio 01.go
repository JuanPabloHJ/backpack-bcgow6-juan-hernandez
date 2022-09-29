package main

import (
	"fmt"
	"os"
)

func ReadArchive(filename string) {
	_, err := os.ReadFile(filename)

	defer fmt.Println("Ejecución finalizada")

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}

func main() {
	ReadArchive("archivo.txt")
}
