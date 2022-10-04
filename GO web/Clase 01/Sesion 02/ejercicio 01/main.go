package main

import (
	"encoding/json"
	"io/ioutil"

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

func Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello Juan Pablo",
	})
}

func GetFile(path string) (file []byte) {
	file, _ = ioutil.ReadFile(path)
	return
}

func GetProducts() (products []product) {
	json.Unmarshal(GetFile("products.JSON"), &products)
	return
}

func GetAll(ctx *gin.Context) {
	ctx.JSON(200, GetProducts())
}

func FilterByColor(ctx *gin.Context) {
	var products []product
	color := ctx.Query("color")

	for _, product := range GetProducts() {
		if product.Color == color {
			products = append(products, product)
		}
	}
	ctx.JSON(200, products)
}

func FilterById(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, product := range GetProducts() {
		if product.Id == id {
			ctx.JSON(200, product)
			return
		}
	}

	ctx.JSON(404, gin.H{"message": "Product not found"})
}

func main() {
	router := gin.Default()

	router.GET("/", Hello)

	products := router.Group("/products")
	{
		products.GET("/", GetAll)
		products.GET("/color", FilterByColor)
		products.GET("/:id", FilterById)
	}

	router.Run()

}
