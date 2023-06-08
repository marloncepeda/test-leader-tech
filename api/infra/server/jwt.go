package server

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func JWTHandler() gin.HandlerFunc {
	authMid, err := InitJWT()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMid.MiddlewareFunc()
}

func InitJWT() (authMiddleware *jwt.GinJWTMiddleware, err error) {
	authMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("&F)J@McQ"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour * 24,
		Authenticator: yourCustomAuthenticatorFunc,
		Unauthorized: func(c *gin.Context, code int, message string) {
			fmt.Println("test ")
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		Authorizator:     yourCustomAuthorizatorFunc,
		PayloadFunc:      yourCustomPayloadFunc,
		IdentityKey:      "id",
		TokenLookup:      "header: Authorization",
		TokenHeadName:    "Bearer",
		TimeFunc:         time.Now,
		SigningAlgorithm: "HS256",
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return
}

func yourCustomAuthenticatorFunc(c *gin.Context) (interface{}, error) {
	// Obtener credenciales de usuario desde el cuerpo de la solicitud
	var credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}
	// Utilizar la función Login para autenticar al usuario
	//user, err := as.Login(AuthModel{Email: credentials.Email, Password: credentials.Password})
	//if err != nil {
	//	return nil, jwt.ErrFailedAuthentication
	//}

	// Retornar el objeto de usuario o una estructura con la información necesaria para generar el token JWT
	return credentials, nil
}

func yourCustomAuthorizatorFunc(data interface{}, c *gin.Context) bool {
	return true
}

func yourCustomPayloadFunc(data interface{}) jwt.MapClaims {

	return jwt.MapClaims{}
}
