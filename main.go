package main

import (
	_ "apiproject/routers"
	"github.com/astaxie/beego"
	"apiproject/controllers"
	"github.com/astaxie/beego/orm"
	"apiproject/bean"
)

func main() {
	// 读取 conf 配置，初始化数据库
	maxIdle,_ := beego.AppConfig.Int("maxIdle")
	maxConn,_ := beego.AppConfig.Int("maxConn")
	datasource := beego.AppConfig.String("datasource")
	orm.RegisterDataBase("default", "mysql", datasource, maxIdle, maxConn)
	orm.RegisterModel(new(bean.User), new(bean.Profile))

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 设置 Controller 方式定义 Error 错误处理函数
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
