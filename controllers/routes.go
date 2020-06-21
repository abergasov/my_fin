package controllers

import (
	"github.com/gin-gonic/gin"
	"my_fin/config"
)

type AppRouter struct {
	GinEngine *gin.Engine
	config    *config.AppConfig
}

func InitRouter(conf *config.AppConfig) *AppRouter {
	if conf.ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	return &AppRouter{
		GinEngine: gin.Default(),
		config:    conf,
	}
}
