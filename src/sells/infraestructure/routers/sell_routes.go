package routers

import (
	"api-hexagonal/src/sells/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterSellRoutes(r *gin.Engine, createSellController * controllers.CreateSellController, getAllSellsController * controllers.GetAllSellsController, updateSellController * controllers.UpdateSellController, deleteSellController * controllers.DeleteSellController) {
	routes := r.Group("/v1/sells")
	{
		routes.POST("/", createSellController.CreateSell)
		routes.GET("/", getAllSellsController.GetAllSells)
		routes.PUT("/:id", updateSellController.EditSell)
		routes.DELETE("/:id", deleteSellController.RemoveSell)
	}
}