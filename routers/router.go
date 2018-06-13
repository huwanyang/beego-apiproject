// @APIVersion 1.0.0
// @Title Beego Project API
// @Description 学习使用 beego 创建一个简单 apiproject 项目，提供 RESTFul API 接口完档，完成基础的 CRUD 操作。。
// @Contact huwanyang168@163.com
package routers

import (
	"apiproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/book",
			beego.NSInclude(
				&controllers.BookController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
