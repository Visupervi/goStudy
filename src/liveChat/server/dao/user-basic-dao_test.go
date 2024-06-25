package dao

import (
	"fmt"
	"testing"
)

//func TestCreateTable(t *testing.T) {
//	CreateTable()
//}

// func TestAddUserBasic(t *testing.T) {
// 	u := &model.UserBasic{}
// 	AddUserBasic(u)

// }

func TestGetUserByName(t *testing.T) {
	user := GetUserByName("TomJerry")
	fmt.Println(user)
}
