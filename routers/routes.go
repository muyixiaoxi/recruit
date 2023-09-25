package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recruit/controllers"
)

func SetupRouter() *gin.Engine {
	// 设置模式
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	v1 := r.Group("/march")

	qq := v1.Group("/qq")
	{
		// 登录
		qq.GET("/login", controllers.Login)

		// 报名
		qq.POST("/signUp", controllers.SignUp)

		// 获取报名信息
		qq.GET("/getMySignUp", controllers.GetMySignUp)

		// 获取所有报名信息
		qq.GET("/getAllSignUp", controllers.GetAllSignUp)

		// 获取未读消息
		qq.GET("/getUnreadMsg", controllers.GetUnreadMsg)
	}

	ding := v1.Group("/ding")
	{
		ding.GET("/getAllSignUp", controllers.GetAllSignUp)
		ding.POST("/sendMsg", controllers.SendMsg)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
