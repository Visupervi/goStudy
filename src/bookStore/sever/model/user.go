package model

// User 实体类
type User struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
