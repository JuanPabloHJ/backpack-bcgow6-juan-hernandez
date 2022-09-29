package main

import "fmt"

const (
	cat     = "cat"
	dog     = "dog"
	hamster = "hamster"
	spider  = "spider"
)

func getCat(quantity float64) float64 {
	return float64(5 * quantity)
}
func getDog(quantity float64) float64 {
	return float64(10 * quantity)
}
func getHamster(quantity float64) float64 {
	return float64(0.250 * quantity)
}
func getSpider(quantity float64) float64 {
	return float64(0.150 * quantity)
}

func Animal(name string) (animal func(float64) float64, err error) {
	switch name {
	case cat:
		animal = getCat
	case dog:
		animal = getDog
	case hamster:
		animal = getHamster
	case spider:
		animal = getSpider
	default:
		err = fmt.Errorf("unknown animal")
	}
	return
}

func main() {
	catFunction, _ := Animal(cat)
	dogFunction, _ := Animal(dog)
	hamsterFunction, _ := Animal(hamster)
	spiderFunction, _ := Animal(spider)

	fmt.Println(catFunction(5))
	fmt.Println(dogFunction(10))
	fmt.Println(hamsterFunction(25))
	fmt.Println(spiderFunction(15))
}
