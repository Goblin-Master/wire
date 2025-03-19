package conf

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"wire/internal/config"
	"wire/internal/controller"
	"wire/internal/middleware"
)

type HttpServer struct {
	Logger         *zap.SugaredLogger
	Engine         *gin.Engine
	UserController *controller.UserController
}

func NewGin(logger *zap.SugaredLogger) *gin.Engine {
	engine := gin.New()
	//加载中间件
	middleware.LoadMiddleware(engine, logger)
	return engine
}

func (h *HttpServer) RunServer() {
	// viper 读取配置文件
	h.Engine.Run(fmt.Sprintf("%s:%s", config.Conf.App.Addr, config.Conf.App.Port))
}
