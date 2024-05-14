package model

import (
	"fmt"
	"goWeb01/db"
)

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *User) AddUser(user *User) (err error) {
	//db.InitDB()
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//db.Db.Query()

	intSmt, err := db.Db.Prepare(sqlStr)

	if err != nil {
		return
	}
	_, err2 := intSmt.Exec(user.UserName, user.Password, user.Email)
	if err2 != nil {
		return
	}
	return
}

func (u *User) AddUser1(user *User) (err error) {
	//db.InitDB()
	sqlStr := "insert into users(username,password, email) values(?,?,?)"
	//db.Db.Query()

	//intSmt, err := db.Db.Prepare(sqlStr)

	//if err != nil {
	//	return
	//}
	//_, err2 := intSmt.Exec(user.UserName, user.Password, user.Email)
	//if err2 != nil {
	//	return
	//}
	//return
	_, err = db.Db.Exec(sqlStr, user.UserName, user.Password, user.Email)
	if err != nil {
		return
	}
	return
}

func (u *User) GetUserById(key int) (*User, error) {
	sqlStr := "select id, username, password, email from users where id = ?"

	row := db.Db.QueryRow(sqlStr, key)

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
	user := &User{
		ID:       id,
		UserName: username,
		Password: password,
		Email:    email,
	}
	return user, nil
}

func (u *User) GetUsers() ([]*User, error) {
	sqlStr := "select id, username, password, email from users"

	rows, err := db.Db.Query(sqlStr)

	if err != nil {
		fmt.Println("查询出错", err)
		return nil, err
	}

	//
	var users []*User
	for rows.Next() {
		var id int
		var username string
		var password string
		var email string

		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			fmt.Println("查询出错", err)
			return nil, err
		}
		user := &User{
			ID:       id,
			UserName: username,
			Password: password,
			Email:    email,
		}

		users = append(users, user)
	}

	return users, nil
}
