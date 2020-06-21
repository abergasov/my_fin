package main

import (
	"log"
	"my_fin/config"
	"my_fin/controllers"
)

func main() {
	router := controllers.InitRouter(config.InitConf())
	router.GinEngine.POST("/api/auth/login", router.Login)
	log.Fatal(router.GinEngine.Run(":8080"))
}
