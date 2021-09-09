package controller

import (
	"fmt"
	"go_web_project/bookstore/dao"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, _ := dao.CheckUsernameAndPassword(username, password)
	if user.ID > 0{
		//用户名和密码正确
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w,"")
	}else {
		//失败
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w,"")
	}
}

func Regist(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user, _ := dao.CheckUsername(username)
	if user.ID > 0{
		//用户名已存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w,"用户名已存在")
	}else {
		//写入数据库中
		dao.SaveUser(username,password,email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w,"用户名或密码不正确")
	}
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取用户名
	username := r.PostFormValue("username")
	fmt.Println("传入的用户名是：", username)

	user, _ := dao.CheckUsername(username)

	if user.ID > 0{
		w.Write([]byte("用户名已存在"))
	}else {
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))

	}
}

