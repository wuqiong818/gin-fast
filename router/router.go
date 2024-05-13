package router

import (
	router_admin "Campusforum/app/admin/router"
	router_im "Campusforum/app/im/router"
	router_other "Campusforum/app/other/router"
	router_post "Campusforum/app/post/router"
	router_user "Campusforum/app/user/router"
	_ "Campusforum/docs"
	"Campusforum/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))                  //自定义中间件
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //swagger文档
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	router_admin.RegisterAdminRoutes(r)

	router_im.RegisterIMRoutes(r)

	router_post.RegisterPostRoutes(r)

	router_other.RegisterOtherRoutes(r)

	router_user.RegisterUserRoutes(r)

	return r
}
