package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type eiData struct {
	Rows            [3]int `json:"rows"`
	Percent         int    `json:"percent"`
	PercentOptional int    `json:"percent_optional"`
}

func (ar *AppRouter) BulkHomePage(c *gin.Context) {
	userID := ar.getUserIDFromRequest(c)
	data, percent, percentOptional := ar.statisticsRepo.RadarCount(userID)
	radarData := &eiData{
		Rows:            data,
		Percent:         percent,
		PercentOptional: percentOptional,
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":       true,
		"ei_radar": radarData,
		"expenses": ar.expenseRepo.GetExpense(userID),
		"per_day":  ar.statisticsRepo.PerDayExp(userID),
	})
}
