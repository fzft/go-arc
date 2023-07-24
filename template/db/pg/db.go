package pg

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB(dataSourceName string) error {
	var err error
	Db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	if err = Db.Ping(); err != nil {
		return err
	}

	return nil
}
