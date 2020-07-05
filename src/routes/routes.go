package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_fin/config"
	"my_fin/src/repository"
	"strconv"
)

type AppRouter struct {
	GinEngine            *gin.Engine
	config               *config.AppConfig
	categoryRepository   *repository.CategoryRepository
	expenseRepository    *repository.ExpenseRepository
	userRepository       *repository.UserRepository
	statisticsRepository *repository.StatisticsRepository
}

type RouterRepoConfig struct {
	CategoryRepository   *repository.CategoryRepository
	ExpenseRepository    *repository.ExpenseRepository
	UserRepository       *repository.UserRepository
	StatisticsRepository *repository.StatisticsRepository
}

func InitRouter(cnf *config.AppConfig, rrC *RouterRepoConfig) *AppRouter {
	if cnf.ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	return &AppRouter{
		GinEngine:            gin.Default(),
		config:               cnf,
		categoryRepository:   rrC.CategoryRepository,
		expenseRepository:    rrC.ExpenseRepository,
		userRepository:       rrC.UserRepository,
		statisticsRepository: rrC.StatisticsRepository,
	}
}

func (ar *AppRouter) getUserIdFromRequest(c *gin.Context) uint64 {
	userId := c.MustGet("user_id")
	uId, err := strconv.Atoi(fmt.Sprintf("%v", userId))
	if err != nil {
		return 0
	}
	return uint64(uId)
}
