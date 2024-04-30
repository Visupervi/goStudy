package controller

import "fmt"

// 为了获取在线用户
var (
	userManager *UserManager
)

type UserManager struct {
	OnlineUsers map[int]*UserController
}

func init() {
	userManager = &UserManager{
		OnlineUsers: make(map[int]*UserController, 1024),
	}
}
func (um *UserManager) AddOnlineUser(up *UserController) {
	um.OnlineUsers[up.UserId] = up
}

func (um *UserManager) DelOnlineUser(userId int) {
	delete(um.OnlineUsers, userId)
}

func (um *UserManager) GetOnlineUsers() map[int]*UserController {
	return um.OnlineUsers
}

func (um *UserManager) GetOnlineUserById(userId int) (up *UserController, err error) {
	//return um.OnlineUsers

	// 如何从map取出一个值， 带检测方式
	up, ok := um.OnlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
