package router_post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(r *gin.Engine) {
	postGroup := r.Group("/api/v1/post")
	{
		// 在这里添加 post 模块的路由
		{
			// 在这里添加 admin 模块的路由
			postGroup.GET("/ping", func(context *gin.Context) { // 测试方法，测试jwt是否生效，不写入swagger，只用于api测试
				context.String(http.StatusOK, "pong,我是admin模块哎")
			})
		}
	}
}
