package controllers

import (
	"apiproject/bean"
	"apiproject/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// 关于 User 相关接口操作
type UserController struct {
	beego.Controller
}

// @Title 增加 user
// @Description 增加一个新的 user 操作
// @Param	body	body	bean.UserProfile	true	"针对 user 的 post 请求 body"
// @Success 200 {int} bean.User.id
// @Failure 403 body 不能为空
// @router / [post]
func (u *UserController) AddUser() {
	var userProfile = bean.UserProfile{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &userProfile)
	if err == nil {
		uid, err := models.AddUser(userProfile)
		if err == nil {
			u.Data["json"] = bean.ResultVO{true, "Add user success", map[string]interface{}{"uid": uid}}
		} else {
			u.Data["json"] = bean.ResultVO{false, err.Error(), ""}
		}
	} else {
		u.Data["json"] = bean.ResultVO{false, err.Error(), ""}
	}
	u.ServeJSON()
}

// @Title 获取 user 列表
// @Description 获取所有的 user 列表
// @Success 200 {object} bean.UserProfile
// @router / [get]
func (u *UserController) GetAllUser() {
	userList := models.GetAllUser()
	u.Data["json"] = bean.ResultVO{true, "", userList}
	u.ServeJSON()
}

// @Title 获取 user
// @Description 根据 uid 获取 user
// @Param	uid	path	int	true "根据 path 的 uid 获取 user 信息"
// @Success 200 {object} bean.UserProfile
// @Failure	403	uid is not int
// @router /:uid [get]
func (u *UserController) GetUser() {
	uid, err := u.GetInt(":uid")
	if err == nil && uid != 0 {
		user, err := models.GetUserById(uid)
		if err == nil {
			u.Data["json"] = bean.ResultVO{true, "", user}
		} else {
			u.Data["json"] = bean.ResultVO{false, err.Error(), ""}
		}
	} else {
		u.Data["json"] = bean.ResultVO{false, err.Error(), ""}
	}
	u.ServeJSON()
}

// @Title 更新 user
// @Description 根据 uid 更新 user & profile
// @Param	uid		path	int		true 	"根据 path 的 uid 更新 user & profile 信息"
// @Param	body	body	bean.UserProfile	true	"针对 user 的 put 请求 body"
// @Success 200 {object} bean.UserProfile
// @Failure	403	uid is not int
// @router /:uid [put]
func (u *UserController) UpdateUser() {
	var userProfile = bean.UserProfile{}
	uid, err := u.GetInt(":uid")
	if err == nil && uid != 0 {
		err := json.Unmarshal(u.Ctx.Input.RequestBody, &userProfile)
		if err == nil {
			err := models.UpdateUser(uid, userProfile)
			if err == nil {
				u.Data["json"] = bean.ResultVO{true, "Update user success", userProfile}
			} else {
				u.Data["json"] = bean.ResultVO{false, err.Error(), ""}
			}
		} else {
			u.Data["json"] = bean.ResultVO{false, err.Error(), ""}
		}
	} else {
		u.Data["json"] = bean.ResultVO{false, err.Error(), ""}
	}
	u.ServeJSON()
}

// @Title 删除 user
// @Description 根据 uid 删除 user，逻辑删除
// @Param	uid	path	int	true "根据 path 的 uid 删除 user 信息"
// @Success 200 {string} delete user success
// @Failure	403	uid is not int
// @router /:uid [delete]
func (u *UserController) DeleteUser() {
	uid, err := u.GetInt(":uid")
	if err == nil && uid != 0 {
		err := models.DeleteUser(uid)
		if err == nil {
			u.Data["json"] = bean.ResultVO{true, "Delete user success", ""}
		} else {
			u.Data["json"] = bean.ResultVO{false, err.Error(), ""}
		}
	} else {
		u.Data["json"] = bean.ResultVO{false, err.Error(), ""}
	}
	u.ServeJSON()
}

// @Title 登录
// @Description 根据用户名和密码验证登录
// @Param	username	query	string	true	"username 用户名"
// @Param	password	query	string	true	"password 密码"
// @Success 200 {string} login success
// @Failure 403 parameter is empty
// @router /login [get,post]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if username != "" && password != "" {
		isOK, userProfile, err := models.Login(username, password)
		if err == nil && isOK {
			u.Data["json"] = bean.ResultVO{true, "Login success.", userProfile}
		} else {
			u.Data["json"] = bean.ResultVO{false, err.Error(), ""}
		}
	} else {
		u.Data["json"] = bean.ResultVO{false, "Parameters required.", ""}
	}
	u.ServeJSON()
}
