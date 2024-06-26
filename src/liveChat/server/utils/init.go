package utils

import (
	"fmt"
	"liveChat/server/db"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

var (
	Pool *redis.Pool
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./src/liveChat/server/config") // /src/liveChat/server/config  /server/config
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config--", viper.Get("mysql"))
}

func InitMysql() {
	db.ConnectDB()
}

// InitRedis 初始化redis
func InitRedis() {
	//var pool *redis.Pool
	Pool = &redis.Pool{
		MaxIdle:     viper.GetInt("redis.maxIdle"),       // 最大空闲连接数
		MaxActive:   viper.GetInt("redis.maxActive"),     // 表示最大的连接数，0表示没有限制
		IdleTimeout: viper.GetDuration("redis.duration"), // 最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", viper.GetString("redis.address"))
		},
	}

}
