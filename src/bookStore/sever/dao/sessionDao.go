package dao

import (
	"bookStore/sever/db"
	"bookStore/sever/model"
	"database/sql"
)

// AddSession 生成一个session
func AddSession(session *model.Session) (res sql.Result, err error) {
	//db.InitDB()
	db, error := db.ConnectDB()
	if error != nil {
		return nil, error
	}
	defer db.Close()
	sqlStr := "insert into sessions(session_id,user_id,username) values(?,?,?)"
	res, err = db.Exec(sqlStr, session.SessionId, session.UserId, session.UserName)

	//fmt.Println("res", res)
	if err != nil {
		return nil, err
	}
	return res, err
}

// GetSessionById 根据sessionId获取session
func GetSessionById(sessionId string) (*model.Session, error) {
	db, error := db.ConnectDB()

	if error != nil {
		return nil, error
	}
	defer db.Close()
	sqlStr := "select session_id, user_id, username from sessions where  session_id = ?"
	row := db.QueryRow(sqlStr, sessionId)
	session := &model.Session{}
	// 这里有个坑，不能直接使用user.，在scan里面不能识别，只能在scan里面使用指针
	err := row.Scan(&session.SessionId, &session.UserId, &session.UserName)
	return session, err
}

// DeleteSession 删除session
func DeleteSession(sessionId string) error {
	db, error := db.ConnectDB()
	if error != nil {
		return error
	}
	defer db.Close()

	sqlStr := "delete from sessions where session_id = ?"
	_, err := db.Exec(sqlStr, sessionId)
	if err != nil {
		return err
	}
	return nil
}
