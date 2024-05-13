package router_admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(r *gin.Engine) {
	adminGroup := r.Group("/api/v1/admin")
	{
		// 在这里添加 admin 模块的路由
		adminGroup.GET("/ping", func(context *gin.Context) { // 测试方法，测试是否生效，不写入swagger，只用于api测试
			context.String(http.StatusOK, "pong,我是admin模块哎")
		})
	}
}
