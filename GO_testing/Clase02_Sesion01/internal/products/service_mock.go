package products

import (
	"errors"
)

type ServiceMock struct {
	db        []Product
	readUsed  bool
	writeUsed bool
}

func (sr *ServiceMock) Read(data interface{}) error {
	sr.readUsed = true
	products, ok := data.(*[]Product)

	if !ok {
		return errors.New("it failed!")
	}

	*products = sr.db
	return nil
}

func (sr *ServiceMock) Write(data interface{}) error {
	sr.writeUsed = true
	products, ok := data.([]Product)

	if !ok {
		return errors.New("it failed!")
	}

	sr.db = products
	return nil
}
