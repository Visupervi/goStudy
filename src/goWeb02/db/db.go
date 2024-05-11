package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

//type DB struct {
//
//}

func InitDB() {
	Db, err = sql.Open("mysql", "root:Wei123456@@tcp(127.0.0.1:3306)/frontEndErrorLogs")

	if err != nil {
		fmt.Println("err=", err)
	}
}
