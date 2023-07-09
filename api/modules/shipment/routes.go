package user

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine, ginFunc gin.HandlerFunc) {

	userRoute := route.Group("/user").Use(ginFunc)
	{
		userRoute.GET("/get/:id", GetUser)
		userRoute.POST("/create/", CreateUser)
		userRoute.PUT("/update/", UpdateUser)
	}
}
