package main

import "fmt"

func main() {
	word := "palabra"

	fmt.Println(len(word))

	for i := 0; i < len(word); i++ {
		fmt.Println(word[i : i+1])
	}
}
