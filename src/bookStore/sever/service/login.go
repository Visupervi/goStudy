package service

import "bookStore/sever/model"

func Login(username string, pwd string) (*model.User, error) {
	return CheckUserNameAndPwd(username, pwd)
}
