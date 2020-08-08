package repository

import (
	"database/sql"
	"encoding/json"
	"my_fin/backend/pkg/database"
	"my_fin/backend/pkg/logger"
)

type Category struct {
	ID           int64      `json:"id"`
	Title        string     `json:"title"`
	CategoryType string     `json:"cat_type"`
	Sub          []Category `json:"sub"` // child categories
}

type CategoryRepository struct {
	db *database.DBAdapter
}

func InitCategoryRepository(db *database.DBAdapter) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (cr *CategoryRepository) LoadCategories(userID uint64) (uCat, uICat []Category) {
	var catJSON string
	var catInJSON string
	row := cr.db.SelectRow("SELECT categories, categories_incoming FROM user_category WHERE u_id = ?", userID)
	err := row.Scan(&catJSON, &catInJSON)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	euC := json.Unmarshal([]byte(catJSON), &uCat)
	if euC != nil {
		logger.Info("Return empty outgoing category list for user: ", userID)
	}
	euCI := json.Unmarshal([]byte(catInJSON), &uICat)
	if euCI != nil {
		logger.Info("Return empty incoming category list for user: ", userID)
	}
	return
}

func (cr *CategoryRepository) UpdateCategories(userID uint64, payload *[]Category, tableKey string) bool {
	var sqlR string
	if tableKey == "categories" {
		sqlR = "INSERT INTO user_category (u_id, categories) VALUES (?, ?) ON DUPLICATE KEY UPDATE categories = ?"
	} else {
		sqlR = "INSERT INTO user_category (u_id, categories_incoming) VALUES (?, ?) ON DUPLICATE KEY UPDATE categories_incoming = ?"
	}
	str, err := json.Marshal(payload)
	if err != nil {
		return false
	}
	_, e := cr.db.Exec(sqlR, userID, str, str)
	return e == nil
}
