package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"my_fin/backend/pkg/data_provider"
)

type Category struct {
	Id           int64      `json:"id"`
	Title        string     `json:"title"`
	CategoryType string     `json:"cat_type"`
	Sub          []Category `json:"sub"` //child categories
}

type CategoryRepository struct {
	db *data_provider.DBAdapter
}

func InitCategoryRepository(db *data_provider.DBAdapter) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (cr *CategoryRepository) LoadCategories(userId uint64) (uCat []Category, uICat []Category) {
	var catJson string
	var catInJson string
	row := cr.db.SelectRow("SELECT categories, categories_incoming FROM user_category WHERE u_id = ?", userId)
	err := row.Scan(&catJson, &catInJson)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	json.Unmarshal([]byte(catJson), &uCat)
	json.Unmarshal([]byte(catInJson), &uICat)
	return
}

func (cr *CategoryRepository) UpdateCategories(userId uint64, payload *[]Category, tableKey string) bool {
	sqlR := fmt.Sprintf("INSERT INTO user_category (u_id, %s) VALUES (?, ?) ON DUPLICATE KEY UPDATE %s = ?", tableKey, tableKey)
	str, err := json.Marshal(payload)
	if err != nil {
		return false
	}
	_, e := cr.db.Exec(sqlR, userId, str, str)
	if e != nil {
		return false
	}
	return true
}
