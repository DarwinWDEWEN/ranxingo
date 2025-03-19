package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/DarwinWDEWEN/ranxingo/types"
	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.Errorf("Handler Panic: %v", r)
				debug.PrintStack()
				c.JSON(http.StatusOK, types.BizVo{Code: types.Failed, Message: types.ErrorMsg})
				c.Abort()
			}
		}()
		//加载完 defer recover，继续后续接口调用
		c.Next()
	}
}
