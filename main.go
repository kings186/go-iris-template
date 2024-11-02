package main

import (
	"go-iris-template/controller"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// 创建用户相关的路由组
	userParty := app.Party("/users")

	// 添加路由
	userParty.Handle("POST", "/register", new(controller.UserController).Register)
	userParty.Handle("POST", "/login", new(controller.UserController).Login)
	userParty.Handle("GET", "/list", new(controller.UserController).List)

	// 启动服务器
	app.Listen(":8080")
}
