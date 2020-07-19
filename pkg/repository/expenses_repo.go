package repository

import (
	"my_fin/backend/pkg/data_provider"
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
	DebtID      int64  `json:"debt_id"`
	Amount      int64  `json:"amount"`
	Commentary  string `json:"commentary"`
	CreatedAt   int64  `json:"created_at"`
	PaymentDate int64  `json:"payment_date"`
	DebtType    int64  `json:"debt_type"`
	ActiveDebt  int64  `json:"active_debt"`
}

const DebtTaken = 1
const DebtGiven = 0

type ExpenseRepository struct {
	db *data_provider.DBAdapter
}

func InitExpenseRepository(db *data_provider.DBAdapter) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (cr *ExpenseRepository) AddExpense(userID uint64, expense *Expense) bool {
	timeNow := time.Now()
	id := cr.db.InsertQuery("expenses", map[string]interface{}{
		"user_id":    userID,
		"created_at": timeNow.Unix(),
		"category":   expense.Category,
		"amount":     expense.Amount,
		"commentary": expense.Commentary,
		"type":       expense.Incoming,
	})
	return id > 0
}

func (cr *ExpenseRepository) GetExpense(userID uint64) *[]Expense {
	sqlR := "SELECT created_at, category, amount, commentary, type FROM expenses WHERE user_id = ? ORDER BY e_id DESC LIMIT 2000"
	rows, err := cr.db.SelectQuery(sqlR, userID)

	var resp []Expense
	if err != nil {
		return &resp
	}

	if rows != nil {
		defer rows.Close()
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

func (cr *ExpenseRepository) AddDebt(userID uint64, d *Debt) bool {
	timeNow := time.Now()
	id := cr.db.InsertQuery("debts", map[string]interface{}{
		"user_id":     userID,
		"created_at":  timeNow.Unix(),
		"amount":      d.Amount,
		"commentary":  d.Commentary,
		"until_date":  d.PaymentDate,
		"debt_type":   d.DebtType,
		"active_debt": 1,
	})
	return id > 0
}

func (cr *ExpenseRepository) PayDebts(userID uint64, debtID, status int64) bool {
	_, e := cr.db.Exec("UPDATE debts SET active_debt = ? WHERE d_id = ? AND user_id = ?", status, debtID, userID)
	return e == nil
}

func (cr *ExpenseRepository) GetDebts(userID uint64) *[]Debt {
	sqlD := "SELECT d_id, created_at, amount, until_date, commentary, debt_type, active_debt FROM debts d WHERE d.user_id = ?"
	rows, err := cr.db.SelectQuery(sqlD, userID)

	var resp []Debt
	if err != nil {
		return &resp
	}

	if rows != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var d Debt

		var tp []uint8
		var ad []uint8
		errS := rows.Scan(&d.DebtID, &d.CreatedAt, &d.Amount, &d.PaymentDate, &d.Commentary, &tp, &ad)
		if errS != nil {
			continue
		}
		d.DebtType = int64(tp[0])
		d.ActiveDebt = int64(ad[0])
		resp = append(resp, d)
	}
	return &resp
}
