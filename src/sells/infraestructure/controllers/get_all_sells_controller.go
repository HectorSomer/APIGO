package controllers

import (
	"fmt"
	"net/http"
	"api-hexagonal/src/sells/application"

	"github.com/gin-gonic/gin"
)

type GetAllSellsController struct {
	useCase *application.GetAllSellsUseCase
}

func NewGetAllSellsController(useCase *application.GetAllSellsUseCase) *GetAllSellsController {
    return &GetAllSellsController{useCase: useCase}
}

func (gpc *GetAllSellsController) GetAllSells(c *gin.Context) {
    sells, err := gpc.useCase.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    var response []gin.H
    for _, sell := range *sells {
        productResponse := gin.H{
            "type": "sell",
            "idSell":   sell.ID,
            "attributes": gin.H{
              "concept":  sell.Concept,
              "date": sell.Date,
              "total_price": sell.Total_Price,
            },
        }
        response = append(response, productResponse)
    }

    if len(*sells) > 0 {
        c.JSON(http.StatusOK, gin.H{"data": response})
    } else {
        fmt.Println("Products:", len(*sells))
      
        c.JSON(http.StatusOK, gin.H{
            "data_size": len(*sells),
            "message": "No se encontraron productos",
        })
    }
}
