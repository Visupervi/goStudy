package redisConfig

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	Pool *redis.Pool
)

//InitRedis 初始化redis
func InitRedis(address string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	//var pool *redis.Pool
	Pool = &redis.Pool{
		MaxIdle:     maxIdle,     // 最大空闲连接数
		MaxActive:   maxActive,   // 表示最大的连接数，0表示没有限制
		IdleTimeout: idleTimeout, // 最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
	//return

}
