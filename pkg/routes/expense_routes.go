package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"my_fin/backend/pkg/repository"
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

func (ar *AppRouter) GetDebts(c *gin.Context) {
	userId := ar.getUserIdFromRequest(c)
	c.JSON(http.StatusOK, gin.H{"ok": true, "debts": ar.expenseRepo.GetDebts(userId)})
}

func (ar *AppRouter) AddDebt(c *gin.Context) {
	var d repository.Debt
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&d)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	userId := ar.getUserIdFromRequest(c)
	result := ar.expenseRepo.AddDebt(userId, &d)
	if result {
		c.JSON(http.StatusOK, gin.H{"ok": result, "debts": ar.expenseRepo.GetDebts(userId)})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"ok": result})
}

func (ar *AppRouter) PayDebt(c *gin.Context) {
	userId := ar.getUserIdFromRequest(c)
	var d struct {
		DebtId     int64 `json:"debt_id"`
		DebtActive int64 `json:"debt_active"`
	}
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&d)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	if ar.expenseRepo.PayDebts(userId, d.DebtId, d.DebtActive) {
		c.JSON(http.StatusOK, gin.H{"ok": true, "debts": ar.expenseRepo.GetDebts(userId)})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
}
