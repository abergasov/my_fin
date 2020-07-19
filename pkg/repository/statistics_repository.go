package repository

import (
	"my_fin/backend/pkg/database"
	"time"
)

const ExpenseIncomingType = "I"
const ExpenseOutgoingType = "E"
const ExpenseMandatoryOutgoingType = "Em"

type StatisticsRepository struct {
	db *database.DBAdapter
}

func InitStatisticsRepository(db *database.DBAdapter) *StatisticsRepository {
	return &StatisticsRepository{db: db}
}

func (sr *StatisticsRepository) RadarCount(userID uint64) (data [3]int, percent, percentMandatory int) {
	sqlQ := "SELECT SUM(amount), type FROM expenses WHERE user_id = ? AND created_at BETWEEN ? AND ? GROUP BY type"
	now := time.Now()
	rows, err := sr.db.SelectQuery(sqlQ, userID, now.Unix()-30*86400, now.Unix())
	if err != nil {
		return [3]int{}, 0, 0
	}
	if rows != nil {
		defer rows.Close()
	}

	incomingSum := 0
	outgoingSum := 0
	outgoingSumMandatory := 0
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
		if typeE == ExpenseMandatoryOutgoingType {
			outgoingSumMandatory = amount
		}
	}
	percent = int(float64(outgoingSum) / float64(incomingSum) * 100)
	percentMandatory = int(float64(outgoingSumMandatory) / float64(incomingSum) * 100)
	incomingSum = incomingSum - outgoingSum - outgoingSumMandatory
	if incomingSum < 0 {
		incomingSum = 0
		percent = -percent
		percentMandatory = -percentMandatory
	}
	return [3]int{incomingSum, outgoingSum, outgoingSumMandatory}, percent, percentMandatory
}

func (sr *StatisticsRepository) GroupedByCategory(userID uint64) interface{} {
	return ""
}
