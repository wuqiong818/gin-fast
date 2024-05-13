package router_user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/api/v1/user")
	{
		// 在这里添加 user 模块的路由
		{
			userGroup.GET("/ping", func(context *gin.Context) { // 测试方法，测试jwt是否生效，不写入swagger，只用于api测试
				context.String(http.StatusOK, "pong,我是user模块哎")
			})
		}
	}
}
