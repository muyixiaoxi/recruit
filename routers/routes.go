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

	admin := v1.Group("/admin")
	admin.POST("/login", controllers.UserLogin)

	admin.GET("/InterviewRecord", controllers.InterviewRecord)
	admin.Use(middlewares.JWTAuthMiddleware())
	{
		// 面试记录
		interview := admin.Group("interviewRecord")
		{
			interview.GET("/query", controllers.GetInterviewRecord)
		}

		// 通知
		{
			// 根据id获取通知详细信息
			admin.GET("/getDetailsMsg", controllers.GetDetailMsg)

			// 发送通知
			admin.POST("/sendMsg", controllers.SendMsg)

			// 保存到草稿箱
			admin.POST("/addDraft", controllers.AddDraft)

			// 获取所有草稿箱
			admin.GET("/getAllDraft", controllers.GetAllDraft)

			// 修改草稿箱
			admin.PUT("/updateDraft", controllers.UpdateDraft)

			// 删除草稿箱
			admin.DELETE("/deleteDraft", controllers.DeleteMsg)

			// 获取所有信息
			admin.GET("/getAllMsg", controllers.UserGetAllMsg)
		}

		// 学生
		{
			// 获取报名信息
			admin.GET("/getAllSignUp", controllers.GetAllSignUp)

			// 筛选
			admin.GET("/screen", controllers.ScreenStudents)

			// 根据专业分组
			admin.GET("/groupBySpecialty", controllers.GroupBySpecialty)

			// 获取详细信息
			admin.GET("/getDetailInfo", controllers.GetDetailInfo)

			// 修改学生状态
			admin.PUT("/updateStudentState", controllers.UpdateStudentsState)

			// 同步已参加宣讲学生状态
			admin.PUT("/updateVisitState", controllers.UpdateVisitState)

			// 同步已面试学生状态
			admin.PUT("/updateInterviewState", controllers.UpdateInterviewState)

		}

		// 安排组
		arrange := admin.Group("arrange")
		{

			// 添加安排组
			arrange.POST("/addArrangeGroup", controllers.AddArrangeGroup)

			// 修改宣讲时间
			arrange.PUT("/updateVisitTime", controllers.UpdateVisitTime)

			// 修改面试时间
			arrange.PUT("/updateInterviewTime", controllers.UpdateInterviewTime)

			// 取消安排组时间
			arrange.DELETE("/cancelTime", controllers.CancelTime)

			// 获取安排组
			arrange.GET("/getArrangeGroup", controllers.GetAllArrangeGroup)

			// 获取安排组菜单
			arrange.GET("/menus", controllers.GetArrangeMenus)

			// 获取安排组详细信息
			arrange.GET("/detail", controllers.GetArrangeDetail)

			// 删除安排组
			arrange.DELETE("/delete", controllers.DeleteArrange)

			// 添加列
			arrange.POST("/col", controllers.AddContentId)

			// 删除列
			arrange.DELETE("/col", controllers.DeleteContentId)
		}

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
