package main

import (
	_ "github.com/owenliang/beego-demo/routers"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/owenliang/beego-demo/models"
)

func init() {
	// 注册driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 注册数据库( 最后2个数字是最大空闲连接和最大数据库连接数量)
	orm.RegisterDataBase("default", "mysql", "root:baidu@123@tcp(127.0.0.1:3306)/beego?charset=utf8&loc=Asia%2fShanghai", 1, 10)

	// 注册model
	orm.RegisterModel(&models.User{})
	// 初始化table
	orm.RunSyncdb("default", false, false)
}

func main() {
	beego.Run()
}
