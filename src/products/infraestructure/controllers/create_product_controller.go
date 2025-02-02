package controllers

import (
	"fmt"
	"net/http"
	"api-hexagonal/src/products/application"
	"api-hexagonal/src/products/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	useCase *application.CreateProductUseCase
}

func NewCreateProductController(useCase *application.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{useCase: useCase}
}

func (pc *CreateProductController) CreateProduct(c *gin.Context) {
    var product entities.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("Create Product")
    createdProduct, err := pc.useCase.Execute(&product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    response := gin.H{
        "data": gin.H{
            "type": "product",
            "idProduct":   createdProduct.ID,
            "attributes": gin.H{
                "name":  createdProduct.Name,
                "price": createdProduct.Price,
                "description": createdProduct.Description,
                "stock": createdProduct.Stock,
            },
        },
    }
    c.JSON(http.StatusCreated, response)
}