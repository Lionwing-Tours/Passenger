package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MySqlDB *sql.DB
)

func init() {
	var err error

	mysqlDb, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/tours-lionwing")
	MySqlDB = mysqlDb
	MySqlDB.SetMaxOpenConns(10)
	MySqlDB.SetMaxIdleConns(0)
	if err != nil {
		panic(err.Error())
	}

}
