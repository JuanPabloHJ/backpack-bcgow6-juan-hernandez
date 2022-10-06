package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var products []product
var id int

type product struct {
	Id           string
	Name, Color  string  `binding:"required"`
	Price        float64 `binding:"required"`
	Stock        int     `binding:"required"`
	Code         string  `binding:"required"`
	Published    bool    `binding:"required"`
	CreationDate string  `binding:"required"`
}

func checkRequiredFields(p product) (err string) {
	var missingFields []string

	if p.Name == "" {
		missingFields = append(missingFields, "Name")
	}
	if p.Color == "" {
		missingFields = append(missingFields, "Color")
	}
	if p.Price == 0 {
		missingFields = append(missingFields, "Price")
	}
	if p.Stock == 0 {
		missingFields = append(missingFields, "Stock")
	}
	if p.Code == "" {
		missingFields = append(missingFields, "Code")
	}
	if p.CreationDate == "" {
		missingFields = append(missingFields, "CreationDate")
	}

	return fmt.Sprintf("El campo %s es requrido", missingFields)
}

func checkRequiredToken(token string) (err error) {
	if token != "666" {
		err = fmt.Errorf("Usuario no autorizado")
	}
	return
}

func AddOneProduct(ctx *gin.Context) {
	var req product
	if err := checkRequiredToken(ctx.GetHeader("token")); err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if err := ctx.ShouldBindJSON(&req); err == nil {
		id++
		req.Id = fmt.Sprintf("%d", id)
		products = append(products, req)

		ctx.JSON(200, "message: success")
	} else {
		ctx.JSON(400, checkRequiredFields(req))
	}

}

func main() {
	router := gin.Default()

	productsRoute := router.Group("/products")
	{
		productsRoute.POST("/", AddOneProduct)
		productsRoute.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, products)
		})
	}

	router.Run()

}
