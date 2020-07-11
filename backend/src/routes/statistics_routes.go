package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ar *AppRouter) IEMonth(c *gin.Context) {
	userId := ar.getUserIdFromRequest(c)
	data, percent, percentOptional := ar.statisticsRepo.RadarCount(userId)
	c.JSON(http.StatusOK, gin.H{"ok": true, "rows": data, "percent": percent, "percent_optional": percentOptional})
}

func (ar *AppRouter) Grouped(c *gin.Context) {
	userId := ar.getUserIdFromRequest(c)
	data := ar.statisticsRepo.GroupedByCategory(userId)
	c.JSON(http.StatusOK, gin.H{"ok": true, "rows": data})
}
