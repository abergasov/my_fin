package routes

import (
	"github.com/gin-gonic/gin"
	"my_fin/config"
	"my_fin/src/repository"
)

type AppRouter struct {
	GinEngine          *gin.Engine
	config             *config.AppConfig
	categoryRepository *repository.CategoryRepository
	expenseRepository  *repository.ExpenseRepository
	userRepository     *repository.UserRepository
}

type RouterRepoConfig struct {
	CategoryRepository *repository.CategoryRepository
	ExpenseRepository  *repository.ExpenseRepository
	UserRepository     *repository.UserRepository
}

func InitRouter(cnf *config.AppConfig, rrC *RouterRepoConfig) *AppRouter {
	if cnf.ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	return &AppRouter{
		GinEngine:          gin.Default(),
		config:             cnf,
		categoryRepository: rrC.CategoryRepository,
		expenseRepository:  rrC.ExpenseRepository,
		userRepository:     rrC.UserRepository,
	}
}
