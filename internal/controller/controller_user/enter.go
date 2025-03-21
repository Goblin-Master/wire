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
	group := engine.Group("/user")
	group.GET("", userController.GetUserInfo)
	return userController
}
func (u *UserController) GetUserInfo(c *gin.Context) {
	u.Logger.Infof("获取用户信息")
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
		c.JSON(200, err)
		return
	}
	c.JSON(200, resp)
}
