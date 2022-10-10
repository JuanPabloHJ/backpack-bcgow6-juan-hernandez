package product

type Service interface {
	GetAll() ([]Product, error)
	Store(id int,
		name string,
		color string,
		price float64,
		stock int,
		code string,
		published bool,
		creationDate string) (Product, error)
	Delete(id int) error
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
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
