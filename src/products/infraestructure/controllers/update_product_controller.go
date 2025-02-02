package controllers

import (
	"api-hexagonal/src/products/application"
	"api-hexagonal/src/products/domain/entities"
    "net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateProductController struct {
	useCase *application.UpdateProductUseCase
}

func NewUpdateProductController(useCase *application.UpdateProductUseCase) *UpdateProductController {
    return &UpdateProductController{useCase: useCase}
}

func (upc *UpdateProductController) UpdateProduct(c *gin.Context) {
	var product entities.UpdateProduct
	idProduct := c.Param("id")

	// Convertir ID
	id, err := strconv.Atoi(idProduct) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La ID no es válida"})
		return
	}

	// Bind JSON solo una vez
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ejecutar la actualización
	productUpdated, err := upc.useCase.Execute(id, product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"detail": err.Error(),
			"type":   "products",
		})
		return
	}

	// Responder con JSON
	response := gin.H{
		"data": gin.H{
			"type": "product",
			"id":   id,
			"attributes": gin.H{
				"name":        productUpdated.Name,
				"price":       productUpdated.Price,
				"description": productUpdated.Description, 
				"stock":       productUpdated.Stock,
			},
		},
	}
	c.JSON(http.StatusOK, response)
}