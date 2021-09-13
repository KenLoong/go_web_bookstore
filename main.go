package main

import (
	"go_web_project/bookstore/controller"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w,"")
}

func main() {

	//设置处理静态资源
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("views/static"))))
	//直接去html页面
	http.Handle("/pages/",http.StripPrefix("/pages/",http.FileServer(http.Dir("views/pages"))))

	http.HandleFunc("/main",IndexHandler)

	http.HandleFunc("/login",controller.Login)
	http.HandleFunc("/logout",controller.Logout)
	http.HandleFunc("/regist",controller.Regist)
	http.HandleFunc("/checkUsername",controller.CheckUserName)

	http.HandleFunc("/getBooks",controller.GetBooks)
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	//根据价格范围查询
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)


	http.ListenAndServe(":8080",nil)
}
