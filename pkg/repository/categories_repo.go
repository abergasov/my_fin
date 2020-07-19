package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"my_fin/backend/pkg/database"
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
	json.Unmarshal([]byte(catJSON), &uCat)
	json.Unmarshal([]byte(catInJSON), &uICat)
	return
}

func (cr *CategoryRepository) UpdateCategories(userID uint64, payload *[]Category, tableKey string) bool {
	sqlR := fmt.Sprintf("INSERT INTO user_category (u_id, %s) VALUES (?, ?) ON DUPLICATE KEY UPDATE %s = ?", tableKey, tableKey)
	str, err := json.Marshal(payload)
	if err != nil {
		return false
	}
	_, e := cr.db.Exec(sqlR, userID, str, str)
	return e == nil
}
