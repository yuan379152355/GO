package main

//package main

//package main

import (
	"fmt"

	//此处注意“_”表示引用mysql函数中init的方法而无需使用函数
	_ "github.com/go-sql-driver/mysql"

	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	//对应id表字段
	Userid int `db:"userid"`
	//对应name表字段
	Username string `db:"username"`
	//对应passwword表字段
	Password string `db:"password"`
}

func main() {
	var cmd string

	for {
		fmt.Println("请输入命令:")
		//fmt.Scan(&cmd)
		cmd = "getall"
		//fmt.Println("你输入的是:",cmd)

		switch cmd {
		case "getall":
			GetAll()
		default:
			fmt.Println("不能识别的命令")
		}

		fmt.Println()
	}

}

func GetAll() {
	//先看看redis里有没有数据
	conn, _ := redis.Dial("tcp", "localhost:6379")
	defer conn.Close()
	reply, err := conn.Do("lrange", "mlist", 0, -1)
	pkeys, _ := redis.Strings(reply, err)
	fmt.Println(pkeys)

	if len(pkeys) > 0 {
		//如果有
		fmt.Println("从redis获得数据")

		// 从redis里直接读取
		for _, key := range pkeys {
			retStrs, _ := redis.Strings(conn.Do("hgetall", key))
			//fmt.Println(retStrs)
			fmt.Printf("{%s %s %s}\n", retStrs[1], retStrs[3], retStrs[5])
		}

	} else {
		//如果没有
		fmt.Println("从mysql获得数据")

		//查询数据库
		db, _ := sqlx.Open("mysql", "root:1@tcp(localhost:3306)/microservice")
		defer db.Close()

		var persons []Person
		db.Select(&persons, "select userid,username,password from my_use")
		fmt.Println(persons)

		//写入redis并且设置过期时间
		for _, p := range persons {
			//将p以hash形式写入redis
			//_, e1 := conn.Do("hmset", p.Id, "userid", p.Name, "age", p.Age, "rmb", p.Rmb)
			_, e1 := conn.Do("hmset", p.Userid, "userid", p.Username, "username", p.Password, "password")
			//将这个hash的key加入mlist
			_, e2 := conn.Do("rpush", "mlist", p.Userid)

			//设置过期时间
			_, e3 := conn.Do("expire", p.Userid, 60)
			_, e4 := conn.Do("expire", "mlist", 60)

			if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
				fmt.Println(p.Username, "写入失败", e1, e2, e3, e4)
			} else {
				fmt.Println(p.Username, "写入成功")
			}
		}
	}
}
