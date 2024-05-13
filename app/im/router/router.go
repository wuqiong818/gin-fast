package router_im

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterIMRoutes(r *gin.Engine) {
	imGroup := r.Group("/api/v1/im")
	{
		// 在这里添加 im 模块的路由
		{
			// 在这里添加 admin 模块的路由
			imGroup.GET("/ping", func(context *gin.Context) { // 测试方法，测试jwt是否生效，不写入swagger，只用于api测试
				context.String(http.StatusOK, "pong,我是im模块哎")
			})
		}
	}
}
