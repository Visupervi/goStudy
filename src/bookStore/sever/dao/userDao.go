package dao

import (
	"bookStore/sever/db"
	"bookStore/sever/model"
	"database/sql"
	"fmt"
)

func AddUser(user *model.User) (res sql.Result, err error) {
	//db.InitDB()
	db, error := db.ConnectDB()
	if error != nil {
		return nil, error
	}

	defer db.Close()
	sqlStr := "insert into users(username,password, email) values(?,?,?)"
	res, err = db.Exec(sqlStr, user.UserName, user.Password, user.Email)

	//fmt.Println("res", res)
	if err != nil {
		return nil, err
	}
	return res, err
}

func GetUserById(key int) (*model.User, error) {
	db, error := db.ConnectDB()

	if error != nil {
		return nil, error
	}

	defer db.Close()
	sqlStr := "select id, username, password, email from users where id = ?"

	row := db.QueryRow(sqlStr, key)

	//

	var id int
	var username string
	var password string
	var email string

	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		fmt.Println("查询出错", err)
		return nil, err
	}
	user := &model.User{
		ID:       id,
		UserName: username,
		Password: password,
		Email:    email,
	}
	return user, nil
}

func GetUsers(userName string, pwd string) (*model.User, error) {
	db, error := db.ConnectDB()

	if error != nil {
		return nil, error
	}
	defer db.Close()
	sqlStr := "select id, username, password, email from users where username = ? and password = ?"
	row := db.QueryRow(sqlStr, userName, pwd)
	user := &model.User{}
	// 这里有个坑，不能直接使用user.，在scan里面不能识别，只能在scan里面使用指针
	err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
	return user, err
}

func GetUserByName(username string) (*model.User, error) {
	db, error := db.ConnectDB()

	if error != nil {
		return nil, error
	}
	defer db.Close()
	sqlStr := "select id, username, password, email from users where username = ?"
	row := db.QueryRow(sqlStr, username)
	//fmt.Println("row", row)
	user := &model.User{}
	// 这里有个坑，不能直接使用user.，在scan里面不能识别，只能在scan里面使用指针
	err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			// 没有找到任何行
			return nil, nil
		}
		// 处理其他类型的错误
		return nil, err
	}

	return user, nil
}
