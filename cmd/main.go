package main

import (
	"my_fin/backend/pkg/config"
	"my_fin/backend/pkg/database"
	"my_fin/backend/pkg/logger"
	"my_fin/backend/pkg/middleware"
	"my_fin/backend/pkg/repository"
	"my_fin/backend/pkg/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	buildTime string = "_dev"
	buildHash string = "_dev"
)

func main() {
	logger.InitLogger("MyFin App", buildHash, buildTime)
	appConf := config.InitConf()
	dbConnection, err := database.InitConnection(appConf)
	if err != nil {
		logger.Fatal("Error db connect", err)
	}

	routerConf := routes.RouterRepoConfig{
		CategoryRepository:   repository.InitCategoryRepository(dbConnection),
		ExpenseRepository:    repository.InitExpenseRepository(dbConnection),
		UserRepository:       repository.InitUserRepository(dbConnection, appConf.JWTKey, appConf.JWTLive),
		StatisticsRepository: repository.InitStatisticsRepository(dbConnection),
		AssetsRepository:     repository.InitAssetsRepository(dbConnection),
	}

	logger.Info("Config ok")
	logger.Info("Init router")

	router := routes.InitRouter(appConf, &routerConf)
	logger.Info("Router ok")

	router.GinEngine.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":         true,
			"build_time": buildTime,
			"build_hash": buildHash,
		})
	})

	router.GinEngine.POST("/api/auth/login", router.Login)
	router.GinEngine.POST("/api/auth/register", router.Register)
	router.GinEngine.POST("/api/auth/logout", router.Logout)
	router.GinEngine.POST("/api/auth/refresh", router.Refresh)

	userData := router.GinEngine.Group("/api/data")
	userData.Use(middleware.AuthMiddleware(routerConf.UserRepository))
	userData.POST("user_category/get", router.UserCategories)
	userData.POST("user_category/update", router.UpdateUserCategories)
	userData.POST("expense/add", router.AddExpense)
	userData.POST("debt/add", router.AddDebt)
	userData.POST("debt/get", router.GetDebts)
	userData.POST("debt/pay", router.PayDebt)
	userData.POST("expense/list", router.GetExpense)

	userData.POST("statistics/list", router.IEMonth)
	userData.POST("statistics/group", router.Grouped)

	logger.Info("Starting server at port 8080")

	errR := router.GinEngine.Run(":8080")
	if errR != nil {
		logger.Fatal("Common server error", errR)
	}
}
