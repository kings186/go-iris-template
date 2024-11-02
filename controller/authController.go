package controller

import (
	"go-iris-template/database"
	"go-iris-template/model"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type AuthController struct{}

// Register 处理用户注册请求
func (c *AuthController) Register(ctx iris.Context) {
	var user model.User
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}

	if err := database.DB.Where("username = ?", user.Username).First(&model.User{}).Error; err == nil {
		ctx.StatusCode(iris.StatusConflict)
		ctx.JSON(iris.Map{"error": "Username already taken"})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to create user"})
		return
	}

	ctx.JSON(iris.Map{"message": "User registered successfully"})
}

// Login 处理用户登录请求
func (c *AuthController) Login(ctx iris.Context) {
	var user model.User
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}

	var databaseUser model.User
	if err := database.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&databaseUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{"error": "Invalid username or password"})
		} else {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(iris.Map{"message": "User logged in successfully"})
}
