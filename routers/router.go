package routers

import (
	"github.com/astaxie/beego"
	"github.com/owenliang/beego-demo/controllers"
	"github.com/astaxie/beego/context"
	"fmt"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/", &controllers.MainController{}),
	)

	beego.AddNamespace(ns)

	beego.InsertFilter("*", beego.AfterExec, func(ctx *context.Context) {
		// 设置502错误码
		fmt.Println("i am here")
		// ctx.Output.SetStatus(502)
	})
}
