package controller

import (
	"go-iris-template/database"
	"go-iris-template/model"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

// UserController 处理用户相关请求
type UserController struct{}

// @Summary 获取所有用户
// @Description 获取系统中所有用户的列表
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} model.User
// @Failure 500 {string} string "内部服务器错误"
// @Router /users [get]
func (c *UserController) GetAllUsers(ctx iris.Context) {
	var users []model.User
	if err := database.DB.Find(&users).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to retrieve users"})
		return
	}
	ctx.JSON(users)
}

// @Summary 获取单个用户
// @Description 根据用户ID获取单个用户的详细信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path uint64 true "用户ID"
// @Success 200 {object} model.User
// @Failure 400 {string} string "无效的用户ID"
// @Failure 404 {string} string "用户未找到"
// @Failure 500 {string} string "内部服务器错误"
// @Router /users/{id} [get]
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

// @Summary 创建新用户
// @Description 创建一个新的用户
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "用户信息"
// @Success 201 {object} model.User
// @Failure 400 {string} string "无效的输入"
// @Failure 500 {string} string "内部服务器错误"
// @Router /users [post]
func (c *UserController) CreateUser(ctx iris.Context) {
	var user model.User
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}
	if err := database.DB.Create(&user).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to create user"})
		return
	}
	ctx.JSON(user)
}

// @Summary 更新用户信息
// @Description 更新指定ID的用户信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path uint64 true "用户ID"
// @Param user body model.User true "更新的用户信息"
// @Success 200 {object} model.User
// @Failure 400 {string} string "无效的输入"
// @Failure 404 {string} string "用户未找到"
// @Failure 500 {string} string "内部服务器错误"
// @Router /users/{id} [put]
func (c *UserController) UpdateUser(ctx iris.Context) {
	var user model.User
	userID := ctx.Params().GetUint64Default("id", 0)
	if userID == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID"})
		return
	}
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}
	if err := database.DB.First(&model.User{}, userID).Updates(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{"error": "User not found"})
		} else {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": "Failed to update user"})
		}
		return
	}
	ctx.JSON(user)
}

// @Summary 删除用户
// @Description 根据ID删除用户
// @Tags users
// @Accept json
// @Produce json
// @Param id path uint64 true "用户ID"
// @Success 200 {string} string "用户删除成功"
// @Failure 400 {string} string "无效的用户ID"
// @Failure 404 {string} string "用户未找到"
// @Failure 500 {string} string "内部服务器错误"
// @Router /users/{id} [delete]
func (c *UserController) DeleteUser(ctx iris.Context) {
	userID := ctx.Params().GetUint64Default("id", 0)
	if userID == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid user ID"})
		return
	}
	if err := database.DB.Delete(&model.User{}, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{"error": "User not found"})
		} else {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": "Failed to delete user"})
		}
		return
	}
	ctx.JSON(iris.Map{"message": "User deleted successfully"})
}
