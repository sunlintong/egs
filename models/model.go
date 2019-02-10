package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//表的设计 操作（增删改查）放Controller里面

type User struct {
	Id   int
	Name string
	Pwd  string
}

func init() {
	//设计数据库基本信息 default(数据库别名) 数据库类型 用户名密码端口号数据库名编码格式
	orm.RegisterDataBase("default", "mysql", "root:gjl19970125@tcp(localhost:3306)/brdb?charset=utf8")
	//映射model数据 等于create table(表名)
	orm.RegisterModel(new(User))
	//生成表 数据库别名 是否强制更新  是否可见
	orm.RunSyncdb("default", false, true)
}
