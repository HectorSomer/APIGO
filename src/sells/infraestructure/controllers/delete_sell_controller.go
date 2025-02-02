package controllers

import (
	"nombre-del-proyecto/src/sells/application"
     "fmt"
	 "strconv"
	 "net/http"
	"github.com/gin-gonic/gin"
)

type DeleteSellController struct {
	useCase *application.DeleteSellUseCase
}

func NewDeleteSellController(useCase *application.DeleteSellUseCase) *DeleteSellController {
    return &DeleteSellController{useCase: useCase}
}

func (dpc *DeleteSellController) RemoveSell (c *gin.Context){
	id := c.Param("id")
    fmt.Println("Removing Sell", id)
    idParse, err := strconv.ParseInt(id, 10, 32)
    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid sell ID"})
    }
	state, err := dpc.useCase.Execute(int(idParse))
	if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if !state {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sell not found"})
		return 
	}else{
		c.JSON(http.StatusOK, 
			gin.H{
				"message": "Sell removed successfully",
				"idSellRemoved": idParse,
			})
        return     
	}
}