package routers

import (
	"net/http"
	"recruit/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 设置模式
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(middlewares.Cors)
	v1 := r.Group("/mosi/v1")

	v1.Use(middlewares.JWTAuthMiddleware())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
