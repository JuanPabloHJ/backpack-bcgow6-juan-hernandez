package main

import (
	"github.com/gin-gonic/gin"

	"Sesion02/cmd/server/handler"
	"Sesion02/internal/product"
)

func main() {
	repo := product.NewRepository()
	service := product.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")

	{
		pr.POST("/", p.Store())
		pr.GET("/", p.GetAll())
		pr.DELETE("/", p.Delete())
	}

	r.Run()

}
