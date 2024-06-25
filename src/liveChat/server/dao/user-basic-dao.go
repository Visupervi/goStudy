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
	//utils.InitMysql()
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

func GetUserByName(name string) *model.UserBasic {
	user := &model.UserBasic{}
	db.DB.Where("name = ?", name).First(user)
	return user
}

func GetUserByNameAndPwd(name string, password string) *model.UserBasic {
	user := &model.UserBasic{}
	db.DB.Where("name = ? and password = ?", name, password).First(user)
	return user
}
func GetUserByPhone(tel string) *model.UserBasic {
	user := &model.UserBasic{}
	db.DB.Where("phone = ?", tel).First(user)
	return user
}

func GetUserByEmail(email string) *model.UserBasic {
	user := &model.UserBasic{}
	db.DB.Where("email = ?", email).First(user)
	return user
}
func DeleteUserById(id uint) {
	db.DB.Delete(&model.UserBasic{}, id)
}

func UpdateUser(user *model.UserBasic) {
	db.DB.Save(user)
}
