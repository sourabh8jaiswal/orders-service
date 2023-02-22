package main

import (
	"order-service/controllers"
	"order-service/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Welcome to Order Service")
	})

	r.POST("/orders", controllers.AddOrder)
	r.POST("/orders/:id", controllers.UpdateOrder)
	r.GET("/orders", controllers.ListOrders)

	r.Run(":5050")
}
