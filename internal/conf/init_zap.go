package conf

import "go.uber.org/zap"

func InitZap() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	return logger.Sugar()
}
