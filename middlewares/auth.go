package middlewares

import (
	controllers2 "recruit/controllers"
	"recruit/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件，支持从Authorization头部或URI中获取Token
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带 Token 有三种方式
		// 1. 放在请求头
		authHeader := c.Request.Header.Get("Authorization")
		// 2. 放在请求体
		// requestBodyToken := c.PostForm("token")
		// 3. 放在 URI
		uriToken := c.Query("token")

		// 从这三种方式中选择一种进行验证，这里选择了从 Authorization 头部获取
		if authHeader == "" && uriToken == "" {
			controllers2.ResponseError(c, controllers2.CodeNeedLogin)
			c.Abort()
			return
		}

		// 选择从 Authorization 头部获取 Token
		tokenString := authHeader
		if tokenString == "" {
			tokenString = uriToken
		}

		// 按空格分割
		parts := strings.SplitN(tokenString, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers2.ResponseError(c, controllers2.CodeInvalidToken)
			c.Abort()
			return
		}

		// parts[1]是获取到的 tokenString，我们使用之前定义好的解析 JWT 的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers2.ResponseError(c, controllers2.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将当前请求的 username 信息保存到请求的上下文 c 上
		c.Set(controllers2.CtxUserIDKey, mc.UserID)
		c.Next() // 后续的处理函数可以用过 c.Get("username") 来获取当前请求的用户信息
	}
}
