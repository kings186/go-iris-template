package main

import (
	"go-iris-template/config"
	"go-iris-template/controller"
	"go-iris-template/database"

	"github.com/kataras/iris/v12"

	_ "go-iris-template/docs"

	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
)

// @title Go iris template Swagger API
// @version 1.0
// @description This is server to represent the API information.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v2
func main() {

	// 加载配置文件
	config.LoadConfig()

	// 初始化数据库
	database.Init()

	app := iris.New()

	app.Logger().SetLevel("debug")

	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		swagger.URL("/swagger/swagger.json"),
		swagger.DeepLinking(true),
		swagger.Prefix("/swagger"),
	)

	// Register on http://localhost:8080/swagger
	app.Get("/swagger/*any", swaggerUI)
	// And the wildcard one for index.html, *.js, *.css and e.t.c.
	app.Get("/swagger/{any:path}", swaggerUI)

	// 创建用户相关的路由组
	authParty := app.Party("/v2/auth")
	// 用户相关路由
	authCtrl := controller.AuthController{}
	userCtrl := controller.UserController{}

	// 添加路由
	authParty.Handle("POST", "/register", authCtrl.Register)
	authParty.Handle("POST", "/login", authCtrl.Login)

	userParty := app.Party("/v2/user")

	userParty.Handle("GET", "/", userCtrl.GetAllUsers)
	userParty.Handle("GET", "/{id:uint64}", userCtrl.GetUserByID)
	userParty.Handle("POST", "/", userCtrl.CreateUser)
	userParty.Handle("PUT", "/{id:uint64}", userCtrl.UpdateUser)
	userParty.Handle("DELETE", "/{id:uint64}", userCtrl.DeleteUser)

	// 启动服务器
	app.Listen(":8080")
}
