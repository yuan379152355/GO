package tables

import (
	"fmt"
	//此处注意“_”表示引用mysql函数中init的方法而无需使用函数
	_ "github.com/go-sql-driver/mysql"

	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

type SY_user struct {

	//对应id表字段
	Userid int `db:"userid"`
	//对应name表字段
	Username string `db:"username"`
	//对应passwword表字段
	Password string `db:"password"`
}

func Insert()  {
	
}