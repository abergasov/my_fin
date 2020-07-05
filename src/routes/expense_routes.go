package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"my_fin/src/repository"
	"net/http"
)

func (ar *AppRouter) AddExpense(c *gin.Context) {
	var e repository.Expense
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	userId := ar.getUserIdFromRequest(c)
	c.JSON(http.StatusOK, gin.H{"ok": ar.expenseRepo.AddExpense(userId, &e)})
}

func (ar *AppRouter) GetExpense(c *gin.Context) {
	userId := ar.getUserIdFromRequest(c)
	c.JSON(http.StatusOK, gin.H{"ok": true, "rows": ar.expenseRepo.GetExpense(userId)})
}
