package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id int64
	Name string `orm:"size(64);unique"` // varchar(64), 唯一索引
}

func InsertUser(user *User) (err error) {
	// 每次请求的ormer需要新建, 不能并发使用
	orm := orm.NewOrm()

	// 插入
	user.Id, err = orm.Insert(user)
	return
}

func FindUser(user *User) (err error) {
	// 每次请求的ormer需要新建, 不能并发使用
	orm := orm.NewOrm()

	err = orm.Read(user)
	return
}