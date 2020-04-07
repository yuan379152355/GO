package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var Pool redis.Pool

func init() { //init 用于初始化一些参数，先于main执行
	Pool = redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {

	conn := Pool.Get()
	res, err := conn.Do("HSET", "student", "name", "jack")
	fmt.Println(res, err)
	res1, err := redis.String(conn.Do("HGET", "student", "name"))
	fmt.Printf("res:%s,error:%v", res1, err)

}

//0 <nil>
//res:jack,error:<nil>
