package controller

import (
	"fmt"
	"go_web_project/bookstore/dao"
	"go_web_project/bookstore/model"
	"go_web_project/bookstore/utils"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	flag,_ := dao.IsLogin(r)
	if flag{
		//已经登录，去首页
		GetPageBooksByPrice(w,r)
	}else {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		user, _ := dao.CheckUsernameAndPassword(username, password)
		if user.ID > 0{
			//添加Session
			//生成UUID
			uuid := utils.CreateUUID()

			sess := &model.Session{
				SessionID: uuid,
				UserName: user.Username,
				UserID: user.ID,
			}
			dao.AddSession(sess)
			//响应cookie
			cookie := http.Cookie{
				Name: "user",
				Value: uuid,
				HttpOnly: true,
			}
			http.SetCookie(w,&cookie)

			//用户名和密码正确
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w,user)
		}else {
			//失败
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w,"")
		}
	}
}

//Logout //处理用户注销的函数
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//删除数据库中与之对应的Session
		dao.DeleteSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		//将修改之后的cookie发送给浏览器
		http.SetCookie(w, cookie)
	}
	//去首页
	GetPageBooksByPrice(w, r)
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

