package model

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	ClientIp      string
	Identity      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOut      time.Time
	IsLoginOut    bool
	DeviceInfo    string
}

func (ub *UserBasic) TableName() string {
	return "user_basic"
}
