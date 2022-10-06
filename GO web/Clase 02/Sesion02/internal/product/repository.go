package product

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
	GetAll() ([]Product, error)
	Store(id int, name string, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	lastID() (int, error)
}

// Streuct repository

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

func NewRepository() Repository {
	return &repository{}
}
