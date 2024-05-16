package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//var (
//	Db  *sql.DB
//	err error
//)

//type DB struct {
//
//}

//func InitDB() {
//	Db, err = sql.Open("mysql", "root:Wei123456@@tcp(127.0.0.1:3306)/bookStore")
//
//	if err != nil {
//		fmt.Println("err=", err)
//	}
//}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Wei123456@@tcp(127.0.0.1:3306)/bookStore")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
