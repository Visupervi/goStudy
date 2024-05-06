package model

import (
	"day35/message"
	"day39/client/model"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//UserDao 用户数据操作逻辑
type UserDao struct {
	pool *redis.Pool
}

func (ud *UserDao) GetUser(userId int, conn redis.Conn) (user *model.User, err error) {

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

	user = &model.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal error", err)
		return
	}
	fmt.Println("user====>", user)
	return
}

func (ud *UserDao) InsertUser(user *model.User, conn redis.Conn) (ur *model.User, err error) {

	_, err = ud.GetUser(user.UserId, conn)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	data, err := json.Marshal(user)

	if err != nil {
		fmt.Println("********InsertUser json.Marshal失败********", err)
		return
	}
	strData := string(data)
	_, err = redis.Int64(conn.Do("hset", "users", user.UserId, strData))

	if err != nil {
		fmt.Println("********注册用户失败********", err)
		err = ERROR_USER_FAIL
		return
	}
	ur = user
	return
}

func DeleteUser(userId int) bool {
	return true
}

func UpdateUser(user message.LoginMsg) bool {
	return true
}
