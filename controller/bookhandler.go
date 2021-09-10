package controller

import (
	"go_web_project/bookstore/dao"
	"go_web_project/bookstore/model"
	"html/template"
	"net/http"
)

//GetPageBooks 获取带分页的图书
//func GetPageBooks(w http.ResponseWriter, r *http.Request) {
//	//获取页码
//	pageNo := r.FormValue("pageNo")
//	if pageNo == "" {
//		pageNo = "1"
//	}
//	//调用bookdao中获取带分页的图书的函数
//	page, _ := dao.GetPageBooks(pageNo)
//	//解析模板文件
//	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//	//执行
//	t.Execute(w, page)
//}

//GetBooks 获取所有图书
func GetBooks(w http.ResponseWriter, r *http.Request) {
	//调用bookdao中获取所有图书的函数
	books, _ := dao.GetBooks()
	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, books)
}

func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	//获取要更新的图书的id
	bookID := r.FormValue("bookId")
	//调用bookdao中获取图书的函数
	book, _ := dao.GetBookByID(bookID)
	if book.ID > 0 {
		//在更新图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, book)
	} else {
		//在添加图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, "")
	}
}

//GetPageBooks 获取带分页的图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用bookdao中获取带分页的图书的函数
	page, _ := dao.GetPageBooks(pageNo)
	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, page)
}


//GetPageBooksByPrice 获取带分页和价格范围的图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	//获取价格范围
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		//调用bookdao中获取带分页的图书的函数
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		//调用bookdao中获取带分页和价格范围的图书的函数
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		//将价格范围设置到page中
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	//调用IsLogin函数判断是否已经登录
	flag, session := dao.IsLogin(r)

	if flag {
		//已经登录，设置page中的IsLogin字段和Username的字段值
		page.IsLogin = true
		page.Username = session.UserName
	}

	//解析模板文件
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, page)
}
