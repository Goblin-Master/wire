package conf

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strings"
)

func InitZap() *zap.SugaredLogger {
	logger, err := zap.NewDevelopment(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return &stackTruncateCore{Core: core}
		}),
	)
	if err != nil {
		log.Fatalf("无法初始化日志记录器: %v", err)
	}
	return logger.Sugar()
}

type stackTruncateCore struct {
	zapcore.Core
}

func (c *stackTruncateCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	// 处理堆栈信息截断
	for i := range fields {
		if fields[i].Key == "stacktrace" {
			// 获取原始堆栈字符串
			stack := fields[i].String
			// 截断为两层（假设每层占两行）
			truncated := truncateStack(stack, 2)
			fields[i].String = truncated
		}
	}
	return c.Core.Write(ent, fields)
}

func truncateStack(stack string, layers int) string {
	lines := strings.Split(stack, "\n")
	maxLines := layers * 2 // 每层包含两行
	if len(lines) > maxLines {
		// 保留指定层数，并添加截断提示
		lines = append(lines[:maxLines], "... stack truncated ...")
	}
	return strings.Join(lines, "\n")
}
