package controllers

import (
	"github.com/astaxie/beego"
	"github.com/owenliang/beego-demo/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
}

func (c *MainController) Get() {
	// 查找Id=1的记录
	user := models.User{Id: 1}

	// 找到记录
	if err := models.FindUser(&user); err == nil {
		// 设置模板数据
		c.Data["user"] = &user
	}

	// 设置模板文件
	c.TplName = "main.html"
}