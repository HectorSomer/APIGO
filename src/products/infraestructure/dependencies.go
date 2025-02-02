package infraestructure

import (
	"fmt"
	"nombre-del-proyecto/src/products/application"

	"nombre-del-proyecto/src/products/infraestructure/controllers"
	"nombre-del-proyecto/src/products/infraestructure/routers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// ! Crear producto
	fmt.Println("Initializing")
	ps := NewMySql()

	createProductUseCase := application.NewCreateProductUseCase(ps)
	createProduct_controller := controllers.NewCreateProductController(createProductUseCase)
	//get
	getProducUseCase := application.NewGetProductUseCase(ps)
    getProduct_Controller := controllers.NewGetProductController(getProducUseCase)
    //put
	updateProductUseCase := application.NewUpdateProductUseCase(ps)
	updateProduct_controller := controllers.NewUpdateProductController(updateProductUseCase)
	//delete
	deleteProductUseCase := application.NewDeleteProductUseCase(ps)
	deleteProduct_controller := controllers.NewDeleteProductController(deleteProductUseCase)
	routers.RegisterProductRoutes(router, createProduct_controller, getProduct_Controller, updateProduct_controller, deleteProduct_controller)

}
