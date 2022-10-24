package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	// arrange
	data := []Product{
		{
			Id:     1,
			Name:   "tomatoe",
			Colour: "red",
			Price:  500,
			Stock:  300,
			Code:   "T",
		},
	}
	updatedProduct := Product{
		Id:     1,
		Name:   "NEW",
		Colour: "NEW",
		Price:  666,
		Stock:  666,
		Code:   "NEW",
	}

	db := ServiceMock{db: data, readUsed: false, writeUsed: false}
	repo := NewRepository(&db)
	service := NewService(repo)

	// act
	result, err := service.Update(
		updatedProduct.Id,
		updatedProduct.Name,
		updatedProduct.Colour,
		updatedProduct.Price,
		updatedProduct.Stock,
		updatedProduct.Code,
		updatedProduct.Published,
	)
	// assert
	assert.True(t, db.readUsed)
	assert.Nil(t, err)
	assert.Equal(t, updatedProduct, result)
}

func TestDelete(t *testing.T) {
	// arrange
	data := []Product{
		{
			Id:     1,
			Name:   "tomatoe",
			Colour: "red",
			Price:  500,
			Stock:  300,
			Code:   "T",
		},
	}

	db := ServiceMock{db: data, readUsed: false, writeUsed: false}
	repo := NewRepository(&db)
	service := NewService(repo)
	expectedResult := []Product{}

	// act
	err := service.Delete(1)
	actualResult := db.db

	// assert
	assert.Nil(t, err)

	assert.True(t, db.readUsed)
	assert.True(t, db.writeUsed)

	assert.Equal(t, expectedResult, actualResult)

}
