package controller_user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"wire/internal/common/res"
	"wire/internal/middleware"
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
	group.GET("", middleware.BindQueryMiddleware[type_user.UserInfoRequest], userController.GetUserInfo)
	group.POST("", middleware.BindJsonMiddleware[type_user.UserCreateRequest], userController.CreateUser)
	group.PUT("", middleware.BindJsonMiddleware[type_user.UserUpdateRequest], userController.UpdateUser)
	group.DELETE("", middleware.BindQueryMiddleware[type_user.UserDeleteRequest], userController.DeleteUser)
	return userController
}
func (u *UserController) GetUserInfo(c *gin.Context) {
	req := middleware.GetBind[type_user.UserInfoRequest](c)
	u.Logger.Infof("获取用户信息: %v", req)
	resp, err := u.UserService.GetUserInfo(req)
	if err != nil {
		u.Logger.Errorf("获取用户信息失败: %v", err)
	}
	res.Response(c, resp, err)
}

func (u *UserController) CreateUser(c *gin.Context) {
	req := middleware.GetBind[type_user.UserCreateRequest](c)
	u.Logger.Infof("创建用户: %v", req)
	resp, err := u.UserService.CreateUser(req)
	if err != nil {
		u.Logger.Errorf("创建用户失败: %v", err)
	}
	res.Response(c, resp, err)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	req := middleware.GetBind[type_user.UserUpdateRequest](c)
	u.Logger.Infof("更新用户: %v", req)
	err := u.UserService.UpdateUser(req)
	if err != nil {
		u.Logger.Errorf("更新用户失败: %v", err)
	}
	res.Response(c, nil, err)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	req := middleware.GetBind[type_user.UserDeleteRequest](c)
	u.Logger.Infof("删除用户: %v", req)
	err := u.UserService.DeleteUser(req)
	if err != nil {
		u.Logger.Errorf("删除用户失败: %v", err)
	}
	res.Response(c, nil, err)
}
