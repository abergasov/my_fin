package repository

import (
	"my_fin/src/data_provider"
)

type Expense struct {
	Category   int64   `json:"cat"`
	Amount     float64 `json:"amount"`
	Incoming   string  `json:"incoming"`
	Commentary string  `json:"commentary"`
}

type ExpenseRepository struct {
	db *data_provider.DBAdapter
}

func InitExpenseRepository(db *data_provider.DBAdapter) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (cr *ExpenseRepository) AddExpense(userId int64, expense *Expense) bool {
	id := cr.db.InsertQuery("expenses", map[string]interface{}{
		"user_id":    userId,
		"category":   expense.Category,
		"amount":     expense.Amount,
		"commentary": expense.Commentary,
		"type":       expense.Incoming,
	})
	return id > 0
}
