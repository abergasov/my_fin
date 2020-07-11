package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_fin/config"
	"my_fin/src/repository"
	"strconv"
)

type AppRouter struct {
	GinEngine      *gin.Engine
	config         *config.AppConfig
	categoryRepo   *repository.CategoryRepository
	expenseRepo    *repository.ExpenseRepository
	userRepo       *repository.UserRepository
	statisticsRepo *repository.StatisticsRepository
	assetsRepo     *repository.AssetsRepository
}

type RouterRepoConfig struct {
	CategoryRepository   *repository.CategoryRepository
	ExpenseRepository    *repository.ExpenseRepository
	UserRepository       *repository.UserRepository
	StatisticsRepository *repository.StatisticsRepository
	AssetsRepository     *repository.AssetsRepository
}

func InitRouter(cnf *config.AppConfig, rrC *RouterRepoConfig) *AppRouter {
	if cnf.ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	return &AppRouter{
		GinEngine:      gin.Default(),
		config:         cnf,
		categoryRepo:   rrC.CategoryRepository,
		expenseRepo:    rrC.ExpenseRepository,
		userRepo:       rrC.UserRepository,
		statisticsRepo: rrC.StatisticsRepository,
		assetsRepo:     rrC.AssetsRepository,
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
