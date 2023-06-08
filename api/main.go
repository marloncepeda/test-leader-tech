package main

import (
	"api/infra/server"
	"api/modules/auth"
	"api/modules/user"
	"fmt"
	"log"
)

func main() {
	fmt.Println("init backend testapi /")

	serverHttp := server.NewServer()
	ginEngine := serverHttp.GinEngine()

	user.Routes(ginEngine, server.JWTHandler())
	auth.Routes(ginEngine, server.JWTHandler())

	if err := serverHttp.Start(); err != nil {
		log.Fatal(err)
	}
}
