package main

import (
	"os"

	"github.com/JuanPabloHJ/backpack-bcgow6-juan-hernandez/tree/main/GO_testing/Clase02_Sesion01/cmd/server/handler"
	"github.com/JuanPabloHJ/backpack-bcgow6-juan-hernandez/tree/main/GO_testing/Clase02_Sesion01/docs"
	"github.com/JuanPabloHJ/backpack-bcgow6-juan-hernandez/tree/main/GO_testing/Clase02_Sesion01/internal/products"
	"github.com/JuanPabloHJ/backpack-bcgow6-juan-hernandez/tree/main/GO_testing/Clase02_Sesion01/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//IMPLEMENTAR MIDDLEWARE PARA TOKEN

// @title GoWeb MELI Bootcamp API
// @version 1.0
// @description This API handle products in a file
func main() {
	//leemos el archivo de variables de entorno
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db := store.NewStore(store.FileType, "./products.json")

	repository := products.NewRepository(db)
	service := products.NewService(repository)
	p := handler.NewProductHandler(service)

	r := gin.Default()

	//generamos el endpoint para consultar la documentacion
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	pr.GET("/", p.GetAll())
	pr.POST("/", p.Store())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateNameAndPrice())
	pr.DELETE("/:id", p.Delete())

	r.Run()
}
