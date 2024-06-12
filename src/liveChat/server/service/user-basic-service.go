package service

import (
	"liveChat/server/dao"
	"liveChat/server/model"

	"gorm.io/gorm"
)

func CreateTable() {
	dao.CreateTable()
}

func InsertUserToTable(user *model.UserBasic) *gorm.DB {
	return dao.AddUserBasic(user)
}

func DeleteUserById(id uint) {
	dao.DeleteUserById(id)
}
