package controller

import (
	"go-iris-template/database"
	"go-iris-template/model"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

// AuthController 处理认证相关请求
type AuthController struct{}

// @Summary 注册新用户
// @Description 创建一个新用户
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.User true "用户信息"
// @Success 200 {object} model.User
// @Failure 400 {string} string "无效的输入"
// @Failure 409 {string} string "用户名已被占用"
// @Failure 500 {string} string "内部服务器错误"
// @Router /auth/register [post]
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

// @Summary 用户登录
// @Description 用户登录以获取访问令牌
// @Tags auth
// @Accept json
// @Produce json
// @Param login body model.User true "用户登录信息"
// @Success 200 {string} string "登录成功"
// @Failure 400 {string} string "无效的输入"
// @Failure 401 {string} string "用户名或密码错误"
// @Failure 500 {string} string "内部服务器错误"
// @Router /auth/login [post]
func (c *AuthController) Login(ctx iris.Context) {
	var user model.User
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}

	var dbUser model.User
	if err := database.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&dbUser).Error; err != nil {
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
