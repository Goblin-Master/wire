package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"runtime/debug"
	"strings"
)

// 自定义Recovery（保留2层堆栈）
func recovery(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := truncateGinStack(debug.Stack(), 2) // 关键步骤
				logger.Error("Panic recovered",
					zap.String("path", c.Request.URL.Path),
					zap.Any("error", err),
					zap.String("stack", stack), // 用截断后的堆栈
				)
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}

func truncateGinStack(fullStack []byte, keepLayers int) string {
	lines := strings.Split(string(fullStack), "\n")
	var (
		result     []string
		foundLayer int
		projectKey = "wire/internal/" // 改为你的项目路径特征
	)

	// 第一阶段：跳过所有框架堆栈
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		// 识别并跳过gin框架堆栈（两行为一组）
		if strings.Contains(line, "github.com/gin-gonic/gin") {
			i++ // 跳过下一行文件路径
			continue
		}
		// 识别并跳过标准库堆栈
		if strings.Contains(line, "net/http") || strings.Contains(line, "runtime/") {
			i++
			continue
		}
		// 识别业务层堆栈
		if strings.Contains(line, projectKey) {
			// 捕获两行（调用方法+文件位置）
			if len(result) == 0 {
				result = append(result, "Business stack:") // 添加标记头
			}
			result = append(result, lines[i], lines[i+1])
			foundLayer++
			i++ // 已处理下一行
			if foundLayer >= keepLayers {
				break
			}
		}
	}

	// 第二阶段：未找到业务堆栈时保留关键信息
	if len(result) == 0 && len(lines) > 6 {
		result = append(result, lines[0:6]...) // 保留前3个有效调用
	}

	// 第三阶段：拼接最终结果
	if len(lines) > len(result)/2*2 {
		result = append(result, "...truncated...")
	}
	return strings.Join(result, "\n")
}
