package controllers

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
)

type DoctorController struct {
	beego.Controller
}

type Login_Doctor struct {
	name string `json:"username"`
	pwd  string `json:"password"`
}

func (d *DoctorController) Get() {
	d.TplName = "login.html"
}
func (d *DoctorController) Post() {
	var req Login_Doctor
	//post请求传的参
	json.Unmarshal(d.Ctx.Input.RequestBody, &req)
	log.Println("收到注册请求，进行逻辑处理", req)

	//1.获取用户名和密码
	// name := Login_Doctor.name
	// pwd := Login_Doctor.pwd

	// user, err := models.QuerryUserByNameAndPwd(name, pwd)
	// if err != nil {
	// 	beego.info(err)
	// }
	//2.连接数据库验证用户名密码是否正确
}
