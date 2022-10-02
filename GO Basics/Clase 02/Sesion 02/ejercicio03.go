package main

import "fmt"

const (
	SMALL  = "small"
	MEDIUM = "medium"
	BIG    = "big"
)

type product struct {
	kind  string
	name  string
	price float64
}

type Product interface {
	GetCost() float64
}

func (p *product) GetCost() (cost float64) {
	switch p.kind {
	case SMALL:
		cost = p.price
	case MEDIUM:
		cost = p.price * 1.03
	case BIG:
		cost = p.price*1.06 + 2500
	}
	return
}

func NewProduct(newProduct product) Product {
	return &newProduct
}

type Store struct {
	products []Product
}

func (s *Store) Total() (total float64) {
	for _, product := range s.products {
		total += product.GetCost()
	}
	return
}

func (s *Store) AddProduct(product Product) {
	s.products = append(s.products, product)
}

type eCommerce interface {
	Total() float64
	AddProduct(Product)
}

func NewStore(store Store) (newStore eCommerce) {
	return &store
}

func main() {
	store := NewStore(Store{})

	store.AddProduct(NewProduct(product{
		kind:  SMALL,
		name:  "product 1",
		price: 1000}))

	store.AddProduct(NewProduct(product{
		kind:  MEDIUM,
		name:  "product 2",
		price: 1000}))

	store.AddProduct(NewProduct(product{
		kind:  BIG,
		name:  "product 3",
		price: 1000}))

	fmt.Println(store.Total())

}
