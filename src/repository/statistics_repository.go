package repository

import (
	"my_fin/src/data_provider"
	"time"
)

const ExpenseIncomingType = "I"
const ExpenseOutgoingType = "E"

type StatisticsRepository struct {
	db *data_provider.DBAdapter
}

func InitStatisticsRepository(db *data_provider.DBAdapter) *StatisticsRepository {
	return &StatisticsRepository{db: db}
}

func (sr *StatisticsRepository) RadarCount(userId uint64) ([2]int, int) {
	sqlQ := "SELECT SUM(amount), type FROM expenses WHERE user_id = ? AND created_at BETWEEN ? AND ? GROUP BY type"
	now := time.Now()
	rows, err := sr.db.SelectQuery(sqlQ, userId, now.Unix()-30*86400, now.Unix())
	if err != nil {
		return [2]int{}, 0
	}
	if rows != nil {
		defer rows.Close()
	}

	incomingSum := 0
	outgoingSum := 0
	for rows.Next() {
		var amount int
		var typeE string

		errS := rows.Scan(&amount, &typeE)
		if errS != nil {
			continue
		}
		if typeE == ExpenseIncomingType {
			incomingSum = amount
		}
		if typeE == ExpenseOutgoingType {
			outgoingSum = amount
		}

	}
	incomingSum = incomingSum - outgoingSum
	percent := (outgoingSum / incomingSum) * 100
	if incomingSum < 0 {
		incomingSum = 0
		percent = -percent
	}
	return [2]int{incomingSum, outgoingSum}, percent
}
