package dao

import (
	"liveChat/server/model"
	"testing"
)

func TestCreateTable(t *testing.T) {
	CreateTable()
}

func TestAddUserBasic(t *testing.T) {
	u := &model.UserBasic{}
	AddUserBasic(u)

}
