package dao

import (
	"github.com/google/uuid"
	db2 "liveChat/server/db"
	"liveChat/server/model"
	"time"
)

func CreateTable() {
	db, _ := db2.ConnectDB()
	// 开始映射
	db.AutoMigrate(&model.UserBasic{})
}

func AddUserBasic() {
	db, _ := db2.ConnectDB()
	// Create
	db.Create(&model.UserBasic{
		Name:          "TomJerry",
		Password:      "123456@",
		Phone:         "15554332761",
		Email:         "admin@outlook.com",
		ClientIp:      "192.168.10.1",
		Identity:      uuid.New().String(),
		ClientPort:    "996",
		LoginTime:     time.Now(),
		HeartbeatTime: time.Now(),
		LoginOut:      time.Now(),
		IsLoginOut:    true,
		DeviceInfo:    uuid.New().String(),
	})
}

func Update(ub *model.UserBasic, k string, v interface{}) {
	db, _ := db2.ConnectDB()
	db.Model(ub).Update(k, v)
}
