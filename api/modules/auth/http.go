package auth

import (
	"api/infra/database"
	"api/pkg/utils"
	"net/http"

	"api/internal/auth"

	"github.com/gin-gonic/gin"
)

func LoginAuth(c *gin.Context) {
	var amodel auth.AuthModel

	authRepo := auth.NewAuthRepository(database.InstanceDB().GetConnect())
	authService := auth.NewAuthService(authRepo)

	if err := c.BindJSON(&amodel); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, nil, err.Error()))
		return
	}

	authLogin, err := authService.Login(amodel)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response(http.StatusOK, nil, err.Error()))
	}

	c.JSON(http.StatusOK, utils.Response(http.StatusOK, authLogin, ""))
}
