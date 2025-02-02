package controllers

import (
	"api-hexagonal/src/sells/application"
	"api-hexagonal/src/sells/domain/entities"
     "strconv"
    "net/http"
	"github.com/gin-gonic/gin"
)

type UpdateSellController struct {
	useCase *application.UpdateSellUseCase
}

func NewUpdateSellController (useCase *application.UpdateSellUseCase) *UpdateSellController {
	return &UpdateSellController{useCase: useCase}
}

func (upc *UpdateSellController) EditSell(c *gin.Context) {
	var sell entities.UpdatedSell
	idSell := c.Param("id")

	// Convertir ID
	id, err := strconv.Atoi(idSell) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La ID no es válida"})
		return
	}
	if err := c.ShouldBindJSON(&sell); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//  actualización
	sellUpdated, err := upc.useCase.Execute(id, sell)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"detail": err.Error(),
			"type":   "error",
		})
		return
	}

	//  JSON
	response := gin.H{
		"data": gin.H{
			"type": "sell",
			"id":   id,
			"attributes": gin.H{
				"concept":     sellUpdated.Concept,
				"date":       sellUpdated.Date,
				"total_price": sellUpdated.TotalPrice, 
			},
		},
	}
	c.JSON(http.StatusOK, response)
}