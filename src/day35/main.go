package main

import (
	"day35/redisConfig"
	"day35/server"
	"time"
)

// redis连接池
func main() {
	//var pool *redis.Pool
	//
	//pool = &redis.Pool{
	//	MaxIdle:     8,   // 最大空闲连接数
	//	MaxActive:   0,   // 表示最大的连接数，0表示没有限制
	//	IdleTimeout: 100, // 最大空闲时间
	//	Dial: func() (redis.Conn, error) {
	//		return redis.Dial("tcp", "localhost:6379")
	//	},
	//}
	//
	//c := pool.Get() // 获取一个连接
	//
	//defer c.Close()
	//
	//_, err := c.Do("set", "name", "TomCat")
	//
	//if err != nil {
	//	fmt.Println("设置失败", err)
	//	return
	//}
	//
	//r, err := redis.String(c.Do("get", "name"))
	//
	//if err != nil {
	//	fmt.Println("读取失败", err)
	//	return
	//}
	//
	//fmt.Println(r)
	//
	//
	//
	//fmt.Println(c)
	redisConfig.InitRedis("localhost:6379", 16, 0, 300*time.Second)
	server.Start()

}
