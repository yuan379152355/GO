package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "root:1@tcp(127.0.0.1:3306)/microservice?charset=utf8")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = db
}

func main() {
	result, err := Db.Exec("INSERT INTO my_use (userid,username,password) VALUES (1004,'yuanshuai','123456')")
	if err != nil {
		fmt.Println("insert failed,error： ", err)
		return
	}
	id, _ := result.LastInsertId()
	fmt.Println("insert id is :", id)
	_, err1 := Db.Exec("update my_use set username = 'yuanshuai01' where userid = 1001")
	if err1 != nil {
		fmt.Println("update failed error:", err1)
	} else {
		fmt.Println("update success!")
	}
	_, err2 := Db.Exec("delete from my_use where userid = 1001 ")
	if err2 != nil {
		fmt.Println("delete error:", err2)
	} else {
		fmt.Println("delete success")
	}

}

//insert id is : 1
//update success!
//delete success
