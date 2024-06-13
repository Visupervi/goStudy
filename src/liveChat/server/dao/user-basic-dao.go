package dao

import (
	"liveChat/server/db"

	//db "liveChat/server/db"
	"liveChat/server/model"

	"gorm.io/gorm"
)

func CreateTable() {
	//db, _ := db2.ConnectDB()
	// 开始映射
	db.DB.AutoMigrate(&model.UserBasic{})
}

func AddUserBasic(user *model.UserBasic) *gorm.DB {
	// Create
	return db.DB.Create(user)
}

func Update(ub *model.UserBasic, k string, v interface{}) {
	//db, _ := db2.ConnectDB()
	db.DB.Model(ub).Update(k, v)
}

func GetUserList() []*model.UserBasic {
	list := make([]*model.UserBasic, 10)
	db.DB.Find(&list)
	return list
}

func DeleteUserById(id uint) {
	db.DB.Delete(&model.UserBasic{}, id)
}

func UpdateUser(user *model.UserBasic) {
	db.DB.Save(user)
}
