package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ar *AppRouter) IEMonth(c *gin.Context) {
	userID := ar.getUserIDFromRequest(c)
	data, percent, percentOptional := ar.statisticsRepo.RadarCount(userID)
	c.JSON(http.StatusOK, gin.H{"ok": true, "rows": data, "percent": percent, "percent_optional": percentOptional})
}

func (ar *AppRouter) Grouped(c *gin.Context) {
	userID := ar.getUserIDFromRequest(c)
	data := ar.statisticsRepo.GroupedByCategory(userID)
	c.JSON(http.StatusOK, gin.H{"ok": true, "rows": data})
}
