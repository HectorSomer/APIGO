package controllers

import (
	"fmt"
	"net/http"
	"api-hexagonal/src/products/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	useCase *application.DeleteProductUseCase
}

func NewDeleteProductController(useCase *application.DeleteProductUseCase) *DeleteProductController {
    return &DeleteProductController{useCase: useCase}
}

func (dpc *DeleteProductController) DeleteProduct (c * gin.Context){
	id := c.Param("id")
    fmt.Println("Delete Product", id)
	idParse, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid product ID"})
        return
    }
    state, err := dpc.useCase.Execute(int(idParse))
    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	if !state {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }else{
		c.JSON(http.StatusOK, 
			gin.H{
				"message": "Product deleted successfully",
			     "idProductDeleted": idParse,
                 "state": state,
			})
	}
}