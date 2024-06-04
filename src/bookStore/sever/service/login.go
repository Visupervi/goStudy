package service

import "bookStore/sever/model"

// Login 登陆
func Login(username string, pwd string) (*model.User, error) {
	return CheckUserNameAndPwd(username, pwd)
}
