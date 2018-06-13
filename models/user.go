package models

import (
	"github.com/astaxie/beego/orm"
	"apiproject/bean"
	"github.com/pkg/errors"
	"fmt"
)

var (
	UserList []bean.UserProfile
)

// 新增 User
func AddUser(up bean.UserProfile) (uid  int64, err  error){
	o := orm.NewOrm()
	o.Begin()
	_, errs := o.Raw("insert into user set username = ?,password = ?,flag = ?", up.Username, up.Password, 0).Exec()
	if errs == nil {
		_, errs = o.Raw("insert into user_profile set user_id = ?,address = ?,age = ?,gender = ?,email = ?",
			uid, up.Address, up.Age, up.Gender, up.Email).Exec()
		if errs == nil {
			o.Commit()
			return uid, nil
		} else {
			o.Rollback()
		}
	} else {
		o.Rollback()
	}
	return 0, errors.New("Insert user fail")
}

// 获取所有的 User 列表
func GetAllUser() []bean.UserProfile {
	o := orm.NewOrm()
	num, err := o.Raw("select * from user u inner join user_profile up on up.user_id = u.id where u.flag = 0").QueryRows(&UserList)
	if err == nil && num > 0 {
		for _, u := range UserList {
			fmt.Printf("UserProfile: %v\n", u)
		}
	}
	return UserList
}

// 根据 id 查询 User 信息
func GetUserById(uid int) (user bean.UserProfile, err error) {
	var userProfile bean.UserProfile
	o := orm.NewOrm()
	errs := o.Raw("select * from user u inner join user_profile up on up.user_id = u.id where u.flag = 0 and u.id = ?", uid).QueryRow(&userProfile)
	fmt.Printf("errs: %v, UserProfile: %v\n", errs, userProfile)
	if errs == nil {
		return userProfile, nil
	} else {
		return user, errors.New("User not exists.")
	}
}

// 更新 user 信息
func UpdateUser(uid int, userProfile bean.UserProfile) error {
	o := orm.NewOrm()
	_, err := o.Raw("update user set username = ?,password = ? where id = ?", userProfile.Username, userProfile.Password, uid).Exec()
	if err == nil {
		_, err = o.Raw("update user_profile set address = ?,age = ?,gender = ?,email = ? where user_id = ?",
			userProfile.Address, userProfile.Age, userProfile.Gender, userProfile.Email, uid).Exec()
		if err == nil {
			return nil
		} else {
			return errors.New("Update user fail")
		}
	} else {
		return errors.New("Update user fail")
	}
}

// 删除 user 信息，逻辑删除
func DeleteUser(uid int) (err error){
	o := orm.NewOrm()
	num, errs := o.Raw("update user set flag = 1 where id = ?", uid).Exec()
	if errs == nil {
		num, errs := num.RowsAffected()
		if errs == nil && num > 0 {
			return nil
		} else {
			return errors.New("Delete user fail")
		}
	} else {
		return errors.New("Delete user fail")
	}
}

// user 登录验证
func Login(username, password string) (isOK bool, userProfile bean.UserProfile, err error) {
	o := orm.NewOrm()
	errs := o.Raw("select * from user u inner join user_profile up on up.user_id = u.id where u.username = ? and u.password = ? and u.flag = 0",
		username, password).QueryRow(&userProfile)
	if errs == nil && userProfile.Id != 0 {
		return true, userProfile,nil
	} else {
		return false, userProfile, errors.New("Login fail.")
	}
}
