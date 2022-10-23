package products

import "errors"

type StubRepository struct {
	db     []Product
	isUsed bool
}

func (sr *StubRepository) Read(data interface{}) error {

	sr.isUsed = true
	products, ok := data.(*[]Product)
	if !ok {
		return errors.New("it failed!")
	}

	*products = sr.db
	return nil
}

func (sr *StubRepository) Write(data interface{}) error {
	return nil
}
