package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// tcp

// 用户注册 用户登陆 显示在线列表 群聊 点对点聊天 离线留言

// redis学习 mysql oracle
// 如何在golang中使用redis
// redis的使用
// noSql数据库
// redis, remote dictionary server （远程字典服务器）， redis性能非常高，单机能够达到15w qps
// 通常适合做缓存， 也可以持久化
// 分布式内存数据库，基于内存运行

// redis的基本使用
// redis五大数据类型
// string, hash, list列表 set集合，zset 有序集合
// redis的字符串最大512M
// setex key seconds value 设置超时间

// redis哈希是一个string类型的field和value的映射表。
// hset

// list数据的操作
//lpush

// set集合操作
// 元素无序的，元素值不能重复 sadd smembers srem[删除指定值] sismember[判断值是否是成员]
func main() {
	//server.Start()

	//	添加key-val [set]
	//
	//conn, err := redis.Dial("tcp", "localhost:6379")
	//
	//defer conn.Close()
	//
	//if err != nil {
	//	fmt.Println("链接失败")
	//	return
	//}
	//// 通过go写入数据
	//
	//_, err = conn.Do("set", "name", "tomJerry")
	//
	//if err != nil{
	//	fmt.Println("写入失败")
	//	return
	//
	//}
	//// 取出数据
	//
	//r, err := redis.String(conn.Do("get", "name"))
	//
	//if err != nil{
	//	fmt.Println("读取失败")
	//	return
	//}
	//
	//fmt.Println("读取成功", r)
	//
	////nameString := r.
	//fmt.Println("链接成功。。。。", conn)

	//	如何操作hash

	conn, err := redis.Dial("tcp", "localhost:6379")

	defer conn.Close()

	if err != nil {
		fmt.Println("链接失败")
		return
	}
	// 通过go写入数据

	_, err = conn.Do("HMSet", "person", "name", "zhngsan", "age", 18, "gender", "female")

	if err != nil {
		fmt.Println("写入失败")
		return

	}
	// 取出数据
	//  redigo: unexpected type for String, got type []interface {}
	//  将redis.String修改成redis.Strings
	r, err := redis.Strings(conn.Do("HMGet", "person", "name", "age", "gender"))

	if err != nil {
		fmt.Println("读取失败", err)
		return
	}

	fmt.Println("读取成功", r)

	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}

	//nameString := r.
	fmt.Println("链接成功。。。。")

}
