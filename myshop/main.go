package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
	_ "myshop/routers"
)

func main() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	fmt.Println(err)
	err = orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/myshop?charset=utf8")
	beego.InsertFilter("*",beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowAllOrigins:true,
		AllowMethods: []string{"GET","POST","PUT","PATCH","DELETE","OPTIONS"},
		AllowHeaders:[]string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		ExposeHeaders:[]string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	}))
	beego.Run()
}
