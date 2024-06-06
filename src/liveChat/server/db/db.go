package db

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:Wei123456@@tcp(127.0.0.1:3306)/live_chat?charset=utf8mb4&parseTime=true&loc=Local"))
	if err != nil {
		return nil, err
	}
	//if err := db.Ping(); err != nil {
	//	return nil, err
	//}
	return db, nil
}
