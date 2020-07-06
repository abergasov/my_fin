package repository

import (
	"my_fin/src/data_provider"
	"time"
)

type Expense struct {
	Category   int64   `json:"cat"`
	Amount     float64 `json:"amount"`
	Incoming   string  `json:"incoming"`
	Commentary string  `json:"commentary"`
	CreatedAt  int64   `json:"created_at"`
}

type Debt struct {
	Amount      int64  `json:"amount"`
	Commentary  string `json:"commentary"`
	CreatedAt   int64  `json:"created_at"`
	PaymentDate int64  `json:"payment_date"`
	DebtType    int64  `json:"debt_type"`
}

const DebtTaken = 1
const DebtGiven = 0

type ExpenseRepository struct {
	db *data_provider.DBAdapter
}

func InitExpenseRepository(db *data_provider.DBAdapter) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (cr *ExpenseRepository) AddExpense(userId uint64, expense *Expense) bool {
	timeNow := time.Now()
	id := cr.db.InsertQuery("expenses", map[string]interface{}{
		"user_id":    userId,
		"created_at": timeNow.Unix(),
		"category":   expense.Category,
		"amount":     expense.Amount,
		"commentary": expense.Commentary,
		"type":       expense.Incoming,
	})
	return id > 0
}

func (cr *ExpenseRepository) GetExpense(userId uint64) *[]Expense {
	sqlR := "SELECT created_at, category, amount, commentary, type FROM expenses WHERE user_id = ? ORDER BY e_id DESC LIMIT 2000"
	rows, err := cr.db.SelectQuery(sqlR, userId)

	if rows != nil {
		defer rows.Close()
	}

	var resp []Expense
	if err != nil {
		return &resp
	}
	for rows.Next() {
		var e Expense
		errS := rows.Scan(&e.CreatedAt, &e.Category, &e.Amount, &e.Commentary, &e.Incoming)
		if errS != nil {
			continue
		}
		resp = append(resp, e)
	}
	return &resp
}

func (cr *ExpenseRepository) AddDebt(userId uint64, d *Debt) bool {
	timeNow := time.Now()
	id := cr.db.InsertQuery("debts", map[string]interface{}{
		"user_id":     userId,
		"created_at":  timeNow.Unix(),
		"amount":      d.Amount,
		"commentary":  d.Commentary,
		"until_date":  d.PaymentDate,
		"debt_type":   d.DebtType,
		"active_debt": 1,
	})
	return id > 0
}
