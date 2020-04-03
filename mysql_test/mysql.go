package main

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "fmt"
)

var Db *sqlx.DB

func init()  {
    db, err := sqlx.Open("mysql", "root:1@tcp(127.0.0.1:3307)/test?charset=utf8")
    if err != nil {
        fmt.Println("open mysql failed,", err)
        return
    }
    Db = db
}

func main()  {
    result, err := Db.Exec("INSERT INTO userinfo (username, password, department,email) VALUES (?, ?, ?,?)","wd","123","it","wd@163.com")
    if err != nil{
        fmt.Println("insert failed,error： ", err)
        return
    }
    id,_ := result.LastInsertId()
    fmt.Println("insert id is :",id)
    _, err1 := Db.Exec("update userinfo set username = ? where uid = ?","jack",1)
    if err1 != nil{
        fmt.Println("update failed error:",err1)
    } else {
        fmt.Println("update success!")
    }
    _, err2 := Db.Exec("delete from userinfo where uid = ? ", 1)
    if err2 != nil{
        fmt.Println("delete error:",err2)
    }else{
        fmt.Println("delete success")
    }

}
//insert id is : 1
//update success!
//delete success