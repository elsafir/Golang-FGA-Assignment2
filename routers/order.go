package routes

import (
	"Golang-FGA-Assignment2/controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func NewOrderRouter(db *gorm.DB, app *gin.Engine) {
	controller := controller.NewOrderController(db)

	app.POST("/orders", controller.CreateOrder)
	app.GET("/orders", controller.GetAllOrders)
	app.PUT("/orders/:id", controller.UpdateOrder)
	app.DELETE("/orders/:id", controller.DeleteOrder)
}
