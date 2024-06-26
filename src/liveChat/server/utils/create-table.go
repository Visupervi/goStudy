package utils

import (
	"liveChat/server/db"
	"liveChat/server/model"
)

func CreateTable() {
	//db.DB.Table("user_basic").First(&model.UserBasic{})
	// db.DB.Table("user_session").First(&model.Session{})
	//
	db.DB.AutoMigrate(&model.Session{})
}
