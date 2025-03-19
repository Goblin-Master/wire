package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoadMiddleware(r *gin.Engine, logger *zap.SugaredLogger) {
	r.Use(log(logger))
	r.Use(gin.Recovery())
	r.Use(Cors())
}
