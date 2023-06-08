package user

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine, ginFunc gin.HandlerFunc) {

	userRoute := route.Group("/user")
	{
		userRoute.POST("/create/", CreateUser)
		userRoute.GET("/get/:id", GetUser).Use(ginFunc)
		userRoute.PUT("/update/", UpdateUser).Use(ginFunc)
	}
}
