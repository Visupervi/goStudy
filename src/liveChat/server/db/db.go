package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (err error) {
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")))
	if err != nil {
		return err
	}
	return nil
}
