package main

import (
	"log"
	"my_fin/config"
	"my_fin/src/data_provider"
	"my_fin/src/repository"
	"my_fin/src/routes"
)

func main() {
	appConf := config.InitConf()
	dbConnection, err := data_provider.InitConnection(appConf)
	if err != nil {
		log.Fatal("Error db connect", err.Error())
	}

	router := routes.InitRouter(appConf, &routes.RouterRepoConfig{
		CategoryRepository: repository.InitCategoryRepository(dbConnection),
		ExpenseRepository:  repository.InitExpenseRepository(dbConnection),
		UserRepository:     repository.InitUserRepository(dbConnection, appConf.JWTKey, appConf.JWTLive),
	})

	router.GinEngine.POST("/api/auth/login", router.Login)
	router.GinEngine.POST("/api/auth/register", router.Register)
	router.GinEngine.POST("/api/user_category/get", router.UserCategories)
	router.GinEngine.POST("/api/user_category/update", router.UpdateUserCategories)
	router.GinEngine.POST("/api/expense/add", router.AddExpense)
	router.GinEngine.POST("/api/expense/list", router.GetExpense)
	log.Fatal(router.GinEngine.Run(":8080"))
}
