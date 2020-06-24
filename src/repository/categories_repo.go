package repository

import (
	"database/sql"
	"encoding/json"
	"my_fin/src/data_provider"
)

type Category struct {
	Id    int64      `json:"id"`
	Title string     `json:"title"`
	Sub   []Category `json:"sub"` //child categories
}

type CategoryRepository struct {
	db *data_provider.DBAdapter
}

func InitCategoryRepository(db *data_provider.DBAdapter) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (cr *CategoryRepository) LoadCategories(userId int64) (uCat []Category) {
	var catJson string
	row := cr.db.SelectRow("SELECT categories FROM user_category WHERE u_id = ?", userId)
	err := row.Scan(&catJson)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	json.Unmarshal([]byte(catJson), &uCat)
	return
}

func (cr *CategoryRepository) UpdateCategories(userId int64, payload *[]Category) bool {
	sql := "INSERT INTO user_category (u_id, categories) VALUES (?, ?) ON DUPLICATE KEY UPDATE categories = ?"
	str, err := json.Marshal(payload)
	if err != nil {
		return false
	}
	_, e := cr.db.Exec(sql, userId, str, str)
	if e != nil {
		return false
	}
	return true
}
