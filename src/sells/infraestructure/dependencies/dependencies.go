package dependencies

import (
	"fmt"
	"nombre-del-proyecto/src/sells/infraestructure"
	"nombre-del-proyecto/src/sells/application"
	"nombre-del-proyecto/src/sells/infraestructure/controllers"
	"nombre-del-proyecto/src/sells/infraestructure/routers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// ! Crear venta
	fmt.Println("Initializing")
	ps := infraestructure.NewMySQL()
	createSellUseCase := application.NewCreateSellUseCase(ps)
	createSell_controller := controllers.NewCreateSellController(createSellUseCase)
	// Obtener todas las ventas
	getAllSells:= application.NewGetAllSellsUseCase(ps)
	getAllSells_controller := controllers.NewGetAllSellsController(getAllSells)
	// Actualizar una venta
	updateSell := application.NewUpdateSellUseCase(ps)
	updateSell_controller := controllers.NewUpdateSellController(updateSell)
	// Eliminar una venta
	deleteSell := application.NewDeleteSellUseCase(ps)
	deleteSell_controller := controllers.NewDeleteSellController(deleteSell)
	routers.RegisterSellRoutes(router, createSell_controller, getAllSells_controller, updateSell_controller, deleteSell_controller)
}
