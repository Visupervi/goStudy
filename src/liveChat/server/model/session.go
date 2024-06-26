package model

import "gorm.io/gorm"

// Session 实体类
type Session struct {
	gorm.Model
	UserName  string
	SessionId string
	UserId    int
	Token     string
}

func (session *Session) TableName() string {
	return "user_session"
}
