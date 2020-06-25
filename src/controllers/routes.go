package controllers

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
}

func InitRouter(cnf *config.AppConfig, cR *repository.CategoryRepository, eR *repository.ExpenseRepository) *AppRouter {
	if cnf.ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	return &AppRouter{
		GinEngine:          gin.Default(),
		config:             cnf,
		categoryRepository: cR,
		expenseRepository:  eR,
	}
}
