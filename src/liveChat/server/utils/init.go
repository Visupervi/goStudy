package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"liveChat/server/db"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../server/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("config", viper.Get("mysql"))
}

func InitMysql() {
	db.ConnectDB()
}
