package controller

import (
	"go-iris-template/database"
	"go-iris-template/model"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type UserController struct{}

// GetAllUsers 获取所有用户
func (c *UserController) GetAllUsers(ctx iris.Context) {
	var users []model.User
	if err := database.DB.Find(&users).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to retrieve users"})
		return
	}
	ctx.JSON(users)
}

// GetUserByID 获取单个用户
func (c *UserController) GetUserByID(ctx iris.Context) {
	var user model.User
	userID := ctx.Params().GetUint64Default("id", 0)
	if userID == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID"})
		return
	}
	if err := database.DB.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{"error": "User not found"})
		} else {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": "Failed to retrieve user"})
		}
		return
	}
	ctx.JSON(user)
}
