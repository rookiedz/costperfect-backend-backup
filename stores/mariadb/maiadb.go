package mariadb

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql" //MySQL Driver
)

var db *sql.DB

//Connect ...
func Connect(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
