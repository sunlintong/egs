package controllers

import (
	"hello/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/* 插入操作
	//1.有orm对象
	o := orm.NewOrm()
	//2.有要插入的结构体对象
	user := models.User{}
	//3.对结构体赋值
	user.Name = "111"
	user.Pwd = "123456"
	//4.插入
	//返回值：几行数据 错误信息
	_, err := o.Insert(&user)
	if err != nil { //err有数据，说明插入失败
		beego.Info("插入失败", err)
		return
	}
	*/
	/*查询操作
	//1.orm对象
	o := orm.NewOrm()
	//2.查询的对象
	user := models.User{}
	//3.指定查询对象字段值
	//根据id查询
	//user.Id = 2
	//err := o.Read(&user)
	//
	//根据name来查
	user.Name = "111"
	//4.查询
	err := o.Read(&user, "Name")
	//返回值: 错误信息

	if err != nil { //err不为空，查询失败
		beego.Info("查询失败", err)
		return
	}
	beego.Info("查询成功", user)
	*/

	/*更新操作
	//1.orm对象
	o := orm.NewOrm()
	//2.需要更新的结构体对象
	user := models.User{}
	//3.需要查找更新的数据
	user.Id = 1
	err := o.Read(&user)
	//4.给数据重复赋值
	if err == nil {
		//说明查找成功
		user.Name = "1213232"

		//5.更新操作
		_, err := o.Update(&user)
		if err != nil {
			//出错
			beego.Info("更新失败", err)
			return
		}
	}
	*/
	/*
		//1.orm对象
		o := orm.NewOrm()
		//2.删除的对象

		user := models.User{}
		//3.指定删除哪一条数据

		user.Id = 1
		//4.删除
		_, err := o.Delete(&user)
		if err != nil {
			//删除失败
			beego.Info("删除失败", err)
			return
		}
		c.Data["data"] = "get请求的欢迎注册"
		c.TplName = "test.html"
	*/

	c.TplName = "register.html"
}

func (c *MainController) Post() {
	//1.拿到数据 用户名和密码
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")

	//2.验证是否合法
	if userName == "" || pwd == "" {
		beego.Info("数据不能为空")
		c.Redirect("/register", 302)
		return
	}
	//3。保存到数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd
	_, err := o.Insert(&user)
	if err != nil {
		//说明插入出现问题
		beego.Info("注册失败")
		c.Redirect("/register", 302)
		return
	}
	//4.返回登录页面
	c.Ctx.WriteString("注册成功")
}

func (c *MainController) LoginGet() {
	beego.Info("进入login的get请求")
	c.TplName = "login.html"
}
func (c *MainController) LoginPost() {
	beego.Info("进入login的post请求")

	//1.获取用户名和密码
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	//2.验证用户名和密码是否合法
	if userName == "" || pwd == "" {
		beego.Info("输入不能为空")
		c.TplName = "login.html"
		return
	}

	//3.验证用户名和密码是否正确

	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd

	err := o.Read(&user, "name", "pwd")
	if err != nil {
		beego.Info("用户名或者密码错误")
		c.TplName = "login.html"
		return
	}
	//4.返回结果对应的效果
	c.TplName = "test.html"
}
