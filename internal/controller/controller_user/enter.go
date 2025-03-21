package controller_user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"wire/internal/service/service_user"
	"wire/internal/type/type_user"
)

type UserController struct {
	UserService service_user.IUserService
	Logger      *zap.SugaredLogger
}

func NewUserController(engine *gin.Engine, UserService service_user.IUserService, logger *zap.SugaredLogger) *UserController {
	userController := &UserController{UserService: UserService, Logger: logger}
	group := engine.Group("/api/user")
	group.GET("", userController.GetUserInfo)
	group.POST("", userController.CreateUser)
	group.PUT("", userController.UpdateUser)
	group.DELETE("", userController.DeleteUser)
	return userController
}
func (u *UserController) GetUserInfo(c *gin.Context) {
	var req type_user.UserInfoRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		u.Logger.Errorf("获取用户信息失败: %v", err)
		c.JSON(200, err)
		return
	}
	u.Logger.Infof("获取用户信息: %v", req)
	resp, err := u.UserService.GetUserInfo(req)
	if err != nil {
		u.Logger.Errorf("获取用户信息失败: %v", err)
		c.JSON(200, err.Error())
		return
	}
	c.JSON(200, resp)
}

func (u *UserController) CreateUser(c *gin.Context) {
	var req type_user.UserCreateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		u.Logger.Errorf("创建用户失败: %v", err)
		c.JSON(200, err)
		return
	}
	u.Logger.Infof("创建用户: %v", req)
	resp, err := u.UserService.CreateUser(req)
	if err != nil {
		u.Logger.Errorf("创建用户失败: %v", err)
		c.JSON(200, err)
		return
	}
	c.JSON(200, resp)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var req type_user.UserUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		u.Logger.Errorf("更新用户失败: %v", err)
		c.JSON(200, err)
		return
	}
	u.Logger.Infof("更新用户: %v", req)
	err = u.UserService.UpdateUser(req)
	if err != nil {
		u.Logger.Errorf("更新用户失败: %v", err)
		c.JSON(200, err)
		return
	}
	c.JSON(200, "更新用户成功")
}

func (u *UserController) DeleteUser(c *gin.Context) {
	var req type_user.UserDeleteRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		u.Logger.Errorf("删除用户失败: %v", err)
		c.JSON(200, err)
		return
	}
	u.Logger.Infof("删除用户: %v", req)
	err = u.UserService.DeleteUser(req)
	if err != nil {
		u.Logger.Errorf("删除用户失败: %v", err)
		c.JSON(200, err)
		return
	}
	c.JSON(200, "删除用户成功")
}
