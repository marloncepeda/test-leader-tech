package auth

import (
	"api/infra/server"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine, ginFunc gin.HandlerFunc) {
	a, _ := server.InitJWT()

	authRoute := route.Group("/auth")
	{
		authRoute.POST("/login", NewAuthHandler().LoginAuth)
		//authRoute.POST("/logout/", LogoutAuth)
		authRoute.PUT("/refrest/", a.RefreshHandler)
	}
}
