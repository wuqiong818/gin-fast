package router_other

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterOtherRoutes(r *gin.Engine) {
	otherGroup := r.Group("/api/v1/other")
	{
		// 在这里添加 other 模块的路由
		{
			// 在这里添加 admin 模块的路由
			otherGroup.GET("/ping", func(context *gin.Context) { // 测试方法，测试jwt是否生效，不写入swagger，只用于api测试
				context.String(http.StatusOK, "pong,我是other模块哎")
			})
		}
	}
}
