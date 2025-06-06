package conf

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"wire/internal/config"
	"wire/internal/controller/controller_user"
	"wire/internal/middleware"
)

type HttpServer struct {
	Logger         *zap.SugaredLogger
	Engine         *gin.Engine
	UserController *controller_user.UserController
}

func InitGin() *gin.Engine {
	engine := gin.New()
	//加载中间件
	middleware.LoadMiddleware(engine)
	return engine
}

func (h *HttpServer) RunServer() {
	// viper 读取配置文件
	h.Engine.Run(config.Conf.App.Link())
}
