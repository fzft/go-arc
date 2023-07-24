package sql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB(dataSourceName string) error {
	var err error
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	if err = Db.Ping(); err != nil {
		return err
	}

	return nil
}
