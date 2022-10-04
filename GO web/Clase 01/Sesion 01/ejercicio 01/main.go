package main

import (
	"github.com/gin-gonic/gin"
)

type product struct {
	Id, Name, Color string
	Price           float64
	Stock           int
	Code            string
	Published       bool
	CreationDate    string
}

func GetAll(ctx *gin.Context) {
	products := []product{
		{
			Id:           "1",
			Name:         "Pen",
			Color:        "Red",
			Price:        500,
			CreationDate: "01/02/2022",
			Stock:        50},

		{
			Id:    "2",
			Name:  "Pen",
			Color: "Black",
			Price: 500,
			Stock: 50,
		}}
	ctx.JSON(200, products)
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello Juan Pablo",
		})
	})

	router.GET("/products", GetAll)

	router.Run()
}
