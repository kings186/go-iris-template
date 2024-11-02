package main

import (
	"go-iris-template/config"
	"go-iris-template/controller"
	"go-iris-template/database"

	"github.com/kataras/iris/v12"
)

func main() {

	// 加载配置文件
	config.LoadConfig()

	// 初始化数据库
	database.Init()

	app := iris.New()

	// 创建用户相关的路由组
	authParty := app.Party("/auth")
	// 用户相关路由
	authCtrl := controller.AuthController{}
	userCtrl := controller.UserController{}

	// 添加路由
	authParty.Handle("POST", "/register", authCtrl.Register)
	authParty.Handle("POST", "/login", authCtrl.Login)

	userParty := app.Party("/user")

	userParty.Handle("GET", "/", userCtrl.GetAllUsers)
	userParty.Handle("GET", "/{id:uint64}", userCtrl.GetUserByID)

	// 启动服务器
	app.Listen(":8080")
}
