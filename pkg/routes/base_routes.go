package routes

import (
	"fmt"
	"my_fin/backend/pkg/config"
	"my_fin/backend/pkg/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICountryChecker interface {
	GetMe(ip string) (country, city string)
}

type AppRouter struct {
	GinEngine      *gin.Engine
	config         *config.AppConfig
	categoryRepo   *repository.CategoryRepository
	expenseRepo    *repository.ExpenseRepository
	userRepo       *repository.UserRepository
	statisticsRepo *repository.StatisticsRepository
	assetsRepo     *repository.AssetsRepository
	country        ICountryChecker
}

type RouterRepoConfig struct {
	CategoryRepository   *repository.CategoryRepository
	ExpenseRepository    *repository.ExpenseRepository
	UserRepository       *repository.UserRepository
	StatisticsRepository *repository.StatisticsRepository
	AssetsRepository     *repository.AssetsRepository
	CountryChecker       ICountryChecker
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
		country:        rrC.CountryChecker,
	}
}

func (ar *AppRouter) getUserIDFromRequest(c *gin.Context) uint64 {
	userID := c.MustGet("user_id")
	uID, err := strconv.Atoi(fmt.Sprintf("%v", userID))
	if err != nil {
		return 0
	}
	return uint64(uID)
}

func (ar *AppRouter) GetUserCountry(c *gin.Context) {
	country, city := ar.country.GetMe(c.GetHeader("X-Real-Ip"))
	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"city":    city,
		"country": country,
		"ip":      c.GetHeader("X-Real-Ip"),
	})
}
