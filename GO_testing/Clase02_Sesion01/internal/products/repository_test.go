package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// arrange
	expectedResult := []Product{
		{
			Id:     1,
			Name:   "tomatoe",
			Colour: "red",
			Price:  500,
			Stock:  300,
			Code:   "T",
		},
		{
			Id:     2,
			Name:   "Potatoe",
			Colour: "Brown",
			Price:  500,
			Stock:  300,
			Code:   "P",
		},
	}

	db := StubRepository{db: expectedResult, isUsed: false}
	repo := NewRepository(&db)

	// act
	result, err := repo.GetAll()
	// assert
	assert.True(t, db.isUsed)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestUpdateName(t *testing.T) {
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
		{
			Id:     2,
			Name:   "Potatoe",
			Colour: "Brown",
			Price:  500,
			Stock:  300,
			Code:   "P",
		},
	}
	expectedResult := data[0]
	expectedResult.Name = "NEW NAME"
	expectedResult.Price = 666

	db := StubRepository{db: data, isUsed: false}
	repo := NewRepository(&db)
	// act
	result, err := repo.UpdateNameAndPrice(1, "NEW NAME", 666)

	// assert
	assert.True(t, db.isUsed)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}
