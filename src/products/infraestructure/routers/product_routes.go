package routers

import (
	"api-hexagonal/src/products/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, createProductController *controllers.CreateProductController, getProductController * controllers.GetProductController, updateProductController * controllers.UpdateProductController, deleteProductController * controllers.DeleteProductController) {
	routes := r.Group("/v1/products")
	{
		routes.POST("/", createProductController.CreateProduct)
		routes.GET("/", getProductController.GetAllProducts)
		routes.PUT("/:id", updateProductController.UpdateProduct)
		routes.DELETE("/:id", deleteProductController.DeleteProduct)
	}
}