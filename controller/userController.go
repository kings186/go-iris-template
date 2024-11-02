package controller

import (
	"go-iris-template/model"

	"github.com/kataras/iris/v12"
)

type UserController struct{}

var users = make(map[string]string)

func (c *UserController) Register(ctx iris.Context) {
	var user model.User
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}

	if _, exists := users[user.Username]; exists {
		ctx.StatusCode(iris.StatusConflict)
		ctx.JSON(iris.Map{"error": "Username already taken"})
		return
	}

	users[user.Username] = user.Password
	ctx.JSON(iris.Map{"message": "User registered successfully"})
}

func (c *UserController) Login(ctx iris.Context) {
	var user model.User
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}

	if password, exists := users[user.Username]; exists && password == user.Password {
		ctx.JSON(iris.Map{"message": "User logged in successfully"})
	} else {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Invalid username or password"})
	}
}

func (c *UserController) List(ctx iris.Context) {

	ctx.JSON(users)
}
