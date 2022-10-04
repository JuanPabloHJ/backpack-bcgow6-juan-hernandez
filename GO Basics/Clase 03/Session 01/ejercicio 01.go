package main

import (
	"fmt"
	"os"
)

type Product struct {
	Id       string
	Price    float64
	Quantity int
}

func AddProducts(archivePath string, products ...Product) {
	var text string
	for _, product := range products {
		text += fmt.Sprintf("%v;%.1f;%d\n", product.Id, product.Price, product.Quantity)
	}
	os.WriteFile(archivePath, []byte(text), 0644)
}

func main() {
	// fmt.Println("")
	p1 := Product{Id: "2", Price: 22, Quantity: 33}
	p2 := Product{Id: "3", Price: 33, Quantity: 44}

	AddProducts("archive.csv", p1, p2)

}
