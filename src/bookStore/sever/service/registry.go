package service

import (
	"bookStore/sever/dao"
	"bookStore/sever/model"
)

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
