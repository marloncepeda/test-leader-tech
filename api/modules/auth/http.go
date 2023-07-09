package auth

import (
	"api/infra/database"
	"api/pkg/utils"
	"net/http"

	"api/internal/auth"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *auth.AuthService
}

func NewAuthHandler() *AuthHandler {
	authRepo := auth.NewAuthRepository(database.InstanceDB().GetConnect())
	authService := auth.NewAuthService(authRepo)
	//authRepo and authServices manager errors pending, now returns 1 value in future needs return 2 with errors
	return &AuthHandler{
		authService: authService,
	}
}

func (ah *AuthHandler) LoginAuth(c *gin.Context) {
	var amodel auth.AuthModel

	if err := c.BindJSON(&amodel); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, nil, err.Error()))
		return
	}

	authLogin, err := ah.authService.Login(amodel)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response(http.StatusOK, nil, err.Error()))
	}

	c.JSON(http.StatusOK, utils.Response(http.StatusOK, authLogin, ""))
}
