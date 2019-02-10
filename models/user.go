// package models

// import (
// 	"hello/models"

// 	"github.com/astaxie/beego"
// 	"github.com/astaxie/beego/orm"
// )

// type Users struct {
// 	id   int    `orm:"column(id);pk;auto"`
// 	name string `orm:"column(name)"`
// 	pwd  string `orm:"column(pwd)"`
// }

// func QuerryUserByNameAndPwd(name, pwd string) (user *Users, err error) {

// 	o := orm.NewOrm()
// 	user := models.Users{}
// 	user.name = name
// 	user.pwd = pwd
// 	_, err := o.Read(&user)
// 	if err != nil {
// 		beego.Info("用户名和密码错误")
// 		return nil, err
// 	}
// 	return user, err
// }
