package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ar *AppRouter) IEMonth(c *gin.Context) {
	userId := ar.getUserIdFromRequest(c)
	data, percent := ar.statisticsRepository.RadarCount(userId)
	c.JSON(http.StatusOK, gin.H{"ok": true, "rows": data, "percent": percent})
}
