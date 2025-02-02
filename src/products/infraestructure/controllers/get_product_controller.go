package controllers

import (
	"nombre-del-proyecto/src/products/application"
     "fmt"
	 "net/http"
	"github.com/gin-gonic/gin"
)

type GetProductController struct {
	useCase *application.GetProductUseCase
}

func NewGetProductController(useCase *application.GetProductUseCase) *GetProductController {
    return &GetProductController{useCase: useCase}
}
func (gpc *GetProductController) GetAllProducts(c *gin.Context) {
    products, err := gpc.useCase.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    var response []gin.H
    for _, product := range *products {
        productResponse := gin.H{
            "type": "product",
            "idProduct":   product.ID,
            "attributes": gin.H{
                "name":  product.Name,
                "price": product.Price,
                "description": product.Description,
                "stock": product.Stock,
            },
        }
        response = append(response, productResponse)
    }

    if len(*products) > 0 {
        c.JSON(http.StatusOK, gin.H{"data": response})
    } else {
        fmt.Println("Products:", len(*products))
      
        c.JSON(http.StatusOK, gin.H{
            "data_size": len(*products),
            "message": "No se encontraron productos",
        })
    }
}