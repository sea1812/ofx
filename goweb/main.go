package main

import (
	"QuoteServer/models"
	_ "QuoteServer/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	// 自动建表
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	//初始化Cache
	models.InitCache()
	beego.Run()
}
