# 项目总结
本项目主要是来熟悉go_web开发，了解服务的架构逻辑。通过一个简单的crud项目来快速上手go语言。

## 1.mysql
在数据库方面，要自己下载mysql的驱动。但要用 _ 去屏蔽这个包，这样子才能用。
import (
"database/sql"
_ "github.com/go-sql-driver/mysql"  //使用mysql，一定要导入这个包，但要用_隐起来
)

连接数据库，通过database/sql包下的*sql.DB对象去连接数据库。
后面都是靠这个对象与mysql进行交互。
Db,err := sql.Open("mysql","root:root@tcp(localhost:3306)/bookstore0612")


sql语句的使用，分2种方式。
一种是预处理，另一种是直接执行sql语句。

预处理模式：
```go
//写sql语句
sqlStr := "select username,user_id from users where session_id = ?"
inStmt, err := utils.Db.Prepare(sqlStr)
if err != nil {
    return nil, err
}
//执行
row := inStmt.QueryRow(sessID)
utils.Db.Prepare(sqlStr)
//获取字段值，通过scan方法去获取
//扫描数据库中的字段值为Session的字段赋值
row.Scan(&sess.UserName, &sess.UserID)
```

直接执行模式：
```go
sqlStr := "insert into users (username,password,email) values(?,?,?)"
//执行
_,err := utils.Db.Exec(sqlStr,username,password,email)
```

api说明：
Exec执行一次命令（包括查询、删除、更新、插入等），不返回任何执行结果。
Query执行一次查询，返回多行结果（即Rows），一般用于执行select命令。
QueryRow执行一次查询，并期望返回最多一行结果（即Row）。QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。（如：未找到结果）

## http接口
编写http接口有3种方法，本项目使用最简单的一种。

路由映射
```go
//这里实现路由映射
http.HandleFunc("/getBooks",controller.GetBooks)
}
```

//接口方法
```go
//接口方法，2个参数是固定的，一个相应，一个请求
func GetBooks(w http.ResponseWriter, r *http.Request) {
//调用bookdao中获取所有图书的函数
books, _ := dao.GetBooks()
//解析模板文件
t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//执行
t.Execute(w, books)
```

启动服务
```go
http.ListenAndServe(":8080",nil)
```