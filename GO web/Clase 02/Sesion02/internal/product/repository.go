package product

import (
	"errors"
)

// Variables
var products []Product
var lastID int

// Structs
type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stoch"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creationDate"`
}

type Repository interface {
	Delete(id int) error
	GetAll() ([]Product, error)
	lastID() (int, error)
	PatchNamePrice(id int, name string, price float64) (Product, error)
	Put(id int, name string, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	Store(id int, name string, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
}

// Struct repository

type repository struct{}

func (r *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (r *repository) lastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, name string, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	product := Product{
		Id:           id,
		Name:         name,
		Color:        color,
		Price:        price,
		Stock:        stock,
		Code:         code,
		Published:    published,
		CreationDate: creationDate}

	product.Id = lastID + 1
	products = append(products, product)
	lastID = product.Id

	return product, nil
}

func (r *repository) Put(id int, name string, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {

	product := Product{
		Id:           id,
		Name:         name,
		Color:        color,
		Price:        price,
		Stock:        stock,
		Code:         code,
		Published:    published,
		CreationDate: creationDate}

	for i, p := range products {
		if id == p.Id {
			products[i] = product
			return product, nil
		}
	}

	return Product{}, errors.New("Product not found")

}

func (r *repository) Delete(id int) error {

	for i, p := range products {
		if p.Id == id {
			products[i] = products[len(products)-1]
			products = products[:len(products)-1]
			return nil
		}
	}

	return errors.New("product not found")
}

func (r *repository) PatchNamePrice(id int, name string, price float64) (Product, error) {
	for i, p := range products {
		if p.Id == id {
			products[i].Price = price
			products[i].Name = name

			return products[i], nil

		}
	}

	return Product{}, errors.New("Product not found")
}

func NewRepository() Repository {
	return &repository{}
}
