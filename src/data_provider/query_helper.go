package data_provider

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"my_fin/config"
	"strings"
)

type DBAdapter struct {
	db *sql.DB
}

func InitConnection(conf *config.AppConfig) (*DBAdapter, error) {
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?interpolateParams=true", conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)
	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		return nil, err
	}
	return &DBAdapter{db: db}, nil
}

func (d *DBAdapter) SelectQuery(sqlString string, params ...interface{}) (*sql.Rows, error) {
	return d.db.Query(sqlString, params...)
}

func (d *DBAdapter) SelectRow(sqlString string, params ...interface{}) *sql.Row {
	return d.db.QueryRow(sqlString, params...)
}

func (d *DBAdapter) Exec(sqlString string, params ...interface{}) (sql.Result, error) {
	return d.db.Exec(sqlString, params...)
}

func (d *DBAdapter) InsertQuery(table string, params map[string]interface{}) (id int64) {
	if len(params) == 0 {
		return
	}
	var sqlP []string
	var values []interface{}
	var sqlPl []string

	for i, v := range params {
		sqlP = append(sqlP, i)
		values = append(values, v)
		sqlPl = append(sqlPl, "?")
	}

	sqlU := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(sqlP, ","), strings.Join(sqlPl, ","))
	res, err := d.db.Exec(sqlU, values...)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		return
	}
	return id
}
