package model

//User 用户实体类
type User struct {
	UserId   int    `json:"userId"`
	Pwd      string `json:"pwd"`
	UserName string `json:"userName"`
	Status   int    `json:"status"`
}
