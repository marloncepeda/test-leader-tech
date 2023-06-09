package user

import (
	"api/infra/database"
	"api/internal/user"
	"api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	userRepo := user.NewUserRepository(database.InstanceDB().GetConnect())
	userService := user.NewUserService(userRepo)

	id := c.Param("id")

	user, err := userService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Response(http.StatusOK, user, ""))
}

func CreateUser(c *gin.Context) {
	var umodel user.UserModel

	userRepo := user.NewUserRepository(database.InstanceDB().GetConnect())
	userService := user.NewUserService(userRepo)

	if err := c.BindJSON(&umodel); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userCreation, err := userService.CreateUser(umodel)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response(http.StatusOK, userCreation, err.Error()))
	}

	c.JSON(http.StatusOK, utils.Response(http.StatusOK, userCreation, ""))

}

func UpdateUser(c *gin.Context) {
	var umodel user.UserModel

	userRepo := user.NewUserRepository(database.InstanceDB().GetConnect())
	userService := user.NewUserService(userRepo)

	if err := c.BindJSON(&umodel); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userUpdate, _ := userService.UpdateUser(umodel)

	c.JSON(http.StatusOK, utils.Response(http.StatusOK, userUpdate, ""))
}
