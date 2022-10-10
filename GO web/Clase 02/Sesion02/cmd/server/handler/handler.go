package handler

import (
	"Sesion02/internal/product"

	"github.com/gin-gonic/gin"
)

type request struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creationDate"`
}

type Product struct {
	service product.Service
}

func NewProduct(p product.Service) *Product {
	return &Product{
		service: p,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		ps, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, ps)
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.Id, req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) Delete() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		err := c.service.Delete(req.Id)

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(204, nil)

	}
}

func (c *Product) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		complete := true

		if req.Id == 0 {
			complete = false
		}
		if req.Name == "" {
			complete = false
		}
		if req.Color == "" {
			complete = false
		}
		if req.Price == 0 {
			complete = false
		}
		if req.Stock == 0 {
			complete = false
		}
		if req.Code == "" {
			complete = false
		}
		if req.CreationDate == "" {
			complete = false
		}

		if !complete {
			ctx.JSON(401, gin.H{
				"error": "Incomplete fields",
			})
			return
		}

		p, err := c.service.Put(req.Id, req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) PatchNamePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		p, err := c.service.PatchNamePrice(req.Id, req.Name, req.Price)

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, p)
	}
}
