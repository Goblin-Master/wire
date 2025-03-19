package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"wire/internal/service"
)

type UserController struct {
	UserService service.IUserService
	Logger      *zap.SugaredLogger
}

func NewUserController(engine *gin.Engine, UserService service.IUserService, logger *zap.SugaredLogger) *UserController {
	userController := &UserController{UserService: UserService, Logger: logger}
	group := engine.Group("/user")
	group.GET("/info", userController.GetUserInfo)
	return userController
}
func (u *UserController) GetUserInfo(ctx *gin.Context) {
	info := u.UserService.GetUserInfo()
	ctx.JSON(200, info)
}
