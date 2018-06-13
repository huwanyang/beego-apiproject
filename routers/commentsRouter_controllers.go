package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["apiproject/controllers:BookController"] = append(beego.GlobalControllerRouter["apiproject/controllers:BookController"],
		beego.ControllerComments{
			Method:           "PostBook",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["apiproject/controllers:BookController"] = append(beego.GlobalControllerRouter["apiproject/controllers:BookController"],
		beego.ControllerComments{
			Method:           "GetAllBook",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["apiproject/controllers:BookController"] = append(beego.GlobalControllerRouter["apiproject/controllers:BookController"],
		beego.ControllerComments{
			Method:           "GetBook",
			Router:           `/:bid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["apiproject/controllers:BookController"] = append(beego.GlobalControllerRouter["apiproject/controllers:BookController"],
		beego.ControllerComments{
			Method:           "UpdateBook",
			Router:           `/:bid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["apiproject/controllers:BookController"] = append(beego.GlobalControllerRouter["apiproject/controllers:BookController"],
		beego.ControllerComments{
			Method:           "DeleteBook",
			Router:           `/:bid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
		beego.ControllerComments{
			Method:           "AddUser",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAllUser",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetUser",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UpdateUser",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
		beego.ControllerComments{
			Method:           "DeleteUser",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
