package dao

import (
	"go_web_project/bookstore/model"
	"go_web_project/bookstore/utils"
)

//AddSession 向数据库中添加Session
func AddSession(sess *model.Session) error {
	//写sql语句
	sqlStr := "insert into sessions values(?,?,?)"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession 删除数据库中的Session
func DeleteSession(sessID string) error {
	//写sql语句
	sqlStr := "delete from sessions where session_id = ?"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}
