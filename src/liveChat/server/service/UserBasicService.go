package service

import "liveChat/server/dao"

func CreateTable() {
	dao.CreateTable()
}

func InsertUserToTable() {
	dao.AddUserBasic()
}
