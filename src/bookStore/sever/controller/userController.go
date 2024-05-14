package controller

import (
	"bookStore/sever/dao"
	"bookStore/sever/model"
	"fmt"
)

// CheckUserNameAndPwd 查看用户登陆是否正确
func CheckUserNameAndPwd(username string, pwd string) (*model.User, error) {
	return dao.GetUsers(username, pwd)
}

// CheckUserName 查看用户名是否存在
func CheckUserName(username string) (*model.User, error) {

	fmt.Println("CheckUserName username=", username)
	return dao.GetUserByName(username)

}
func Login(username string, pwd string) (*model.User, error) {
	return CheckUserNameAndPwd(username, pwd)
}

func Registry(username string, pwd string, email string) (bool, error) {
	user, err := CheckUserName(username)
	//fmt.Println("user=")
	if err != nil || user != nil {
		return false, err
	}

	u := &model.User{
		UserName: username,
		Password: pwd,
		Email:    email,
	}
	_, err = dao.AddUser(u)
	if err != nil {
		return false, err
	}
	return true, nil
}
