package middleware

import (
	"github.com/gin-gonic/gin"
)

func LoadMiddleware(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Cors())
}
