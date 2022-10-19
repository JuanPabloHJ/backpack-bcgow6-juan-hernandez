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
		pr.DELETE("/:id", p.Delete())
		pr.GET("/", p.GetAll())
		pr.PATCH("/:id", p.PatchNamePrice())
		pr.POST("/", p.Store())
		pr.PUT("/:id", p.Put())

	}

	r.Run()
}
