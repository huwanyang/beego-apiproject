package controllers

import (
	"apiproject/bean"
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (e *ErrorController) Error401() {
	err := bean.Error{401, "Sorry, API Url Unauthorized."}
	e.Data["json"] = err
	e.ServeJSON()
}

func (e *ErrorController) Error403() {
	err := bean.Error{403, "Sorry, API Url Forbidden."}
	e.Data["json"] = err
	e.ServeJSON()
}

func (e *ErrorController) Error404() {
	err := bean.Error{404, "Sorry, API Url Not Found."}
	e.Data["json"] = err
	e.ServeJSON()
}

func (e *ErrorController) Error500() {
	err := bean.Error{500, "Sorry, Internal Server Error."}
	e.Data["json"] = err
	e.ServeJSON()
}

func (e *ErrorController) Error503() {
	err := bean.Error{503, "Sorry, Service Unavailable."}
	e.Data["json"] = err
	e.ServeJSON()
}
