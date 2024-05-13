package middlewares

import (
	"Campusforum/pkg/jwt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		AuthHeader := c.Request.Header.Get("Authorization")
		if AuthHeader == "" {
			log.Println("Authorization header is empty.")
			//controller.ResponseError(c, controller.CodeNoAuth)
			c.Abort()
			return
		}

		parts := strings.SplitN(AuthHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("Invalid Authorization format.")
			//controller.ResponseError(c, controller.CodeInvalidAuthFormat)
			c.Abort()
			return
		}

		token := strings.TrimSpace(parts[1])
		log.Println("Extracted JWT token:", token)

		// 校验token
		mc, err := jwt.ValidateToken(token)
		if err != nil {
			log.Println("Invalid token:", err)
			//controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		c.Set(jwt.ContextKeyUserObj, mc)
		c.Next()
	}
}
