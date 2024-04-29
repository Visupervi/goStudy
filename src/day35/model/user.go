package model

import (
	"day35/redisConfig"
	"day37/client/model"
	"fmt"
)

//User 用户实体类
type User struct {
	UserId   int    `json:"userId"`
	Pwd      string `json:"pwd"`
	UserName string `json:"userName"`
}

func (u *User) Login(uId int, pwd string) (user *User, err error) {

	c := redisConfig.Pool.Get()
	defer c.Close()
	ud := UserDao{
		pool: redisConfig.Pool,
	}

	user, err = ud.GetUser(uId, c)

	//json.Unmarshal([]byte(user))

	if err != nil {
		fmt.Println("登陆失败")
		return
	}

	if user.Pwd != pwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (u *User) Registry(user *model.User) (err error, err1 error) {
	c := redisConfig.Pool.Get()
	defer c.Close()
	ud := UserDao{
		pool: redisConfig.Pool,
	}

	ud.InsertUser(user, c)
	fmt.Println("ud", ud)
	return
}
