package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recruit/controllers"
	"recruit/middlewares"
)

func SetupRouter() *gin.Engine {
	// 设置模式
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(middlewares.Cors())
	v1 := r.Group("/march")

	qq := v1.Group("/qq")
	{
		// 登录
		qq.POST("/login", controllers.Login)

		// 报名
		qq.POST("/signUp", controllers.SignUp)

		// 获取报名信息
		qq.GET("/getMySignUp", controllers.GetMySignUp)

		//获取我的全部消息
		qq.GET("/getMyAllMsg", controllers.GetMyAllMsg)

		// 已读消息
		qq.POST("/readMsg", controllers.AddReadMsg)
	}

	user := v1.Group("/admin")
	user.POST("/login", controllers.UserLogin)

	user.Use(middlewares.JWTAuthMiddleware())
	{
		// 面试记录
		{
			user.GET("/InterviewRecord", controllers.InterviewRecord)

			user.GET("/getInterviewRecord", controllers.GetInterviewRecord)
		}

		// 通知
		{
			// 根据id获取通知详细信息
			user.GET("/getDetailsMsg", controllers.GetDetailMsg)

			// 发送通知
			user.POST("/sendMsg", controllers.SendMsg)

			// 保存到草稿箱
			user.POST("/addDraft", controllers.AddDraft)

			// 获取所有草稿箱
			user.GET("/getAllDraft", controllers.GetAllDraft)

			// 修改草稿箱
			user.PUT("/updateDraft", controllers.UpdateDraft)

			// 删除草稿箱
			user.DELETE("/deleteDraft", controllers.DeleteMsg)

			// 获取所有信息
			user.GET("/getAllMsg", controllers.UserGetAllMsg)
		}

		// 学生
		{
			// 获取报名信息
			user.GET("/getAllSignUp", controllers.GetAllSignUp)

			// 筛选
			user.GET("/screen", controllers.ScreenStudents)

			// 根据专业分组
			user.GET("/groupBySpecialty", controllers.GroupBySpecialty)

			// 获取详细信息
			user.GET("/getDetailInfo", controllers.GetDetailInfo)

			// 修改学生状态
			user.PUT("/updateStudentState", controllers.UpdateStudentsState)

			// 同步已参加宣讲学生状态
			user.PUT("/updateVisitState", controllers.UpdateVisitState)

			// 同步已面试学生状态
			user.PUT("/updateInterviewState", controllers.UpdateInterviewState)

		}

		// 用户
		{

			// 添加安排组
			user.POST("/addArrangeGroup", controllers.AddArrangeGroup)

			// 修改宣讲时间
			user.PUT("/updateVisitTime", controllers.UpdateVisitTime)

			// 修改面试时间
			user.PUT("/updateInterviewTime", controllers.UpdateInterviewTime)

			// 获取安排组
			user.GET("/getArrangeGroup", controllers.GetAllArrangeGroup)

		}

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
