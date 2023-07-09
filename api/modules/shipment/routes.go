package user

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine, ginFunc gin.HandlerFunc) {

	userRoute := route.Group("/user").Use(ginFunc)
	{
		userRoute.GET("/get/:id", NewShipmentHandler().GetUser)
		userRoute.POST("/create/", NewShipmentHandler().CreateUser)
		userRoute.PUT("/update/", NewShipmentHandler().UpdateUser)
	}
}
