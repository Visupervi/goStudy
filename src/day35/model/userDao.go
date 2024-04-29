package model

import (
	"day37/message"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//UserDao 用户数据操作逻辑
type UserDao struct {
	pool *redis.Pool
}

func (ud *UserDao) GetUser(userId int, conn redis.Conn) (user *User, err error) {

	//p := conn.Get() // 获取一个连接
	//defer p.Close()
	res, err := redis.String(conn.Do("hget", "users", userId))

	fmt.Println("res====>", res)
	//fmt.Printf("res====>%T", res)

	if err != nil {
		//fmt.Println("查询出错")
		if err == redis.ErrNil {
			fmt.Println("********该用户不存在********")
			err = ERROR_USER_NOTEXISTS
		}
		//user = User{}

		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal error", err)
		return
	}
	fmt.Println("user====>", user)
	return
}

func InsertUser() (user message.LoginMsg) {
	return
}

func DeleteUser(userId int) bool {
	return true
}

func UpdateUser(user message.LoginMsg) bool {
	return true
}
