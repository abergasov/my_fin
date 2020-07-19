package routes

import (
	"encoding/json"
	"my_fin/backend/pkg/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ar *AppRouter) AddExpense(c *gin.Context) {
	var e repository.Expense
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	userID := ar.getUserIDFromRequest(c)
	c.JSON(http.StatusOK, gin.H{"ok": ar.expenseRepo.AddExpense(userID, &e)})
}

func (ar *AppRouter) GetExpense(c *gin.Context) {
	userID := ar.getUserIDFromRequest(c)
	c.JSON(http.StatusOK, gin.H{"ok": true, "rows": ar.expenseRepo.GetExpense(userID)})
}

func (ar *AppRouter) GetDebts(c *gin.Context) {
	userID := ar.getUserIDFromRequest(c)
	c.JSON(http.StatusOK, gin.H{"ok": true, "debts": ar.expenseRepo.GetDebts(userID)})
}

func (ar *AppRouter) AddDebt(c *gin.Context) {
	var d repository.Debt
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&d)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	userID := ar.getUserIDFromRequest(c)
	result := ar.expenseRepo.AddDebt(userID, &d)
	if result {
		c.JSON(http.StatusOK, gin.H{"ok": result, "debts": ar.expenseRepo.GetDebts(userID)})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"ok": result})
}

func (ar *AppRouter) PayDebt(c *gin.Context) {
	userID := ar.getUserIDFromRequest(c)
	var d struct {
		DebtID     int64 `json:"debt_id"`
		DebtActive int64 `json:"debt_active"`
	}
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&d)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	if ar.expenseRepo.PayDebts(userID, d.DebtID, d.DebtActive) {
		c.JSON(http.StatusOK, gin.H{"ok": true, "debts": ar.expenseRepo.GetDebts(userID)})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
}
