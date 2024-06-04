package service

import (
	"bookStore/sever/dao"
	"bookStore/sever/model"
)

// CheckUserNameAndPwd 查看用户登陆是否正确
func CheckUserNameAndPwd(username string, pwd string) (*model.User, error) {
	return dao.GetUsers(username, pwd)
}

// CheckUserName 查看用户名是否存在
func CheckUserName(username string) (*model.User, error) {
	//fmt.Println("CheckUserName username=", username)
	return dao.GetUserByName(username)

}
