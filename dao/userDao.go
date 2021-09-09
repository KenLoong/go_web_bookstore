package dao

import (
	"fmt"
	"go_web_project/bookstore/model"
	"go_web_project/bookstore/utils"
)

func CheckUsernameAndPassword(username string,password string)(*model.User, error) {
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr, username, password)

	user := &model.User{}
	row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	fmt.Println("row.user = ",user)
	return user,nil
}

func CheckUsername(username string) (*model.User, error) {
	sqlStr := "select id,username,password,email from users where username = ?"
	row := utils.Db.QueryRow(sqlStr, username)

	user := model.User{}
	row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	return &user,nil
}

func SaveUser(username string, password string, email string) error {
	sqlStr := "insert into users (username,password,email) values(?,?,?)"
	//执行
	_,err := utils.Db.Exec(sqlStr,username,password,email)
	if err != nil{
		return err
	}
	return nil
}
