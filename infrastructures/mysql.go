package infrastructures

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySqlDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
