package model

// Session 实体类
type Session struct {
	UserName  string `json:"userName"`
	SessionId string `json:"sessionId"`
	UserId    int    `json:"userId"`
}
