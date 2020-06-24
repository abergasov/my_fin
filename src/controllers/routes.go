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
}

func InitRouter(conf *config.AppConfig, categoryRepository *repository.CategoryRepository) *AppRouter {
	if conf.ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	return &AppRouter{
		GinEngine:          gin.Default(),
		config:             conf,
		categoryRepository: categoryRepository,
	}
}
