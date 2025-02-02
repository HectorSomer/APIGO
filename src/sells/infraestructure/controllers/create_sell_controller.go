package controllers

import (
	"fmt"
	"net/http"
	"api-hexagonal/src/sells/application"
	"api-hexagonal/src/sells/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateSellController struct {
	useCase *application.CreateSellUseCase
}

func NewCreateSellController(useCase *application.CreateSellUseCase) *CreateSellController {
    return &CreateSellController{useCase: useCase}
}
func (pc *CreateSellController) CreateSell(c *gin.Context) {
    var sell entities.Sell
    if err := c.ShouldBindJSON(&sell); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("Create Sell")
    createdSell, err := pc.useCase.Execute(sell)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    response := gin.H{
        "data": gin.H{
            "type": "sell",
            "idSell":   createdSell.ID,
            "attributes": gin.H{
                "concept":  createdSell.Concept,
                "total_price": createdSell.Total_Price,
                "date": createdSell.Date,
            },
        },
    }
    c.JSON(http.StatusCreated, response)
}