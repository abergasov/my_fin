package main

import (
	"log"
	"my_fin/config"
	"my_fin/src/controllers"
	"my_fin/src/data_provider"
	"my_fin/src/repository"
)

func main() {
	appConf := config.InitConf()
	dbConnection, err := data_provider.InitConnection(appConf)
	if err != nil {
		log.Fatal("Error db connect", err.Error())
	}
	categoryRepository := repository.InitCategoryRepository(dbConnection)

	router := controllers.InitRouter(appConf, categoryRepository)
	router.GinEngine.POST("/api/auth/login", router.Login)
	router.GinEngine.POST("/api/user_category/get", router.UserCategories)
	router.GinEngine.POST("/api/user_category/update", router.UpdateUserCategories)
	log.Fatal(router.GinEngine.Run(":8080"))
}
