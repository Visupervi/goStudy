package service

import (
	"bookStore/sever/dao"
	"bookStore/sever/model"
	"database/sql"
)

func AddSession(sess *model.Session) (res sql.Result, err error) {
	return dao.AddSession(sess)
}

func DeleteSession(sessionId string) error {
	return dao.DeleteSession(sessionId)
}
