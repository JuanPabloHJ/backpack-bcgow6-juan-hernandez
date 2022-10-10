package product

type Service interface {
	Delete(id int) error
	GetAll() ([]Product, error)
	PatchNamePrice(id int, name string, price float64) (Product, error)
	Put(id int,
		name string,
		color string,
		price float64,
		stock int,
		code string,
		published bool,
		creationDate string) (Product, error)
	Store(id int,
		name string,
		color string,
		price float64,
		stock int,
		code string,
		published bool,
		creationDate string) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return products, nil

}

func (s *service) Store(id int, name string, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	lastID, err := s.repository.lastID()

	if err != nil {
		return Product{}, err
	}

	lastID++

	product, err := s.repository.Store(id, name, color, price, stock, code, published, creationDate)

	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) Put(id int, name string, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	return s.repository.Put(id, name, color, price, stock, code, published, creationDate)

}

func (s *service) PatchNamePrice(id int, name string, price float64) (Product, error) {
	return s.repository.PatchNamePrice(id, name, price)
}
