package user

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine, ginFunc gin.HandlerFunc) {
	shipmentHandler := NewShipmentHandler()

	userRoute := route.Group("/user").Use(ginFunc)
	{
		userRoute.GET("/get/:id", shipmentHandler.GetUser)
		userRoute.POST("/create/", shipmentHandler.CreateUser)
		userRoute.PUT("/update/", shipmentHandler.UpdateUser)
	}
}
