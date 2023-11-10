package controllers

// 面试、宣讲、参观安排
import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"recruit/models"
	"recruit/service"
)

// GetAllArrangeGroup 获取全部安排组
func GetAllArrangeGroup(c *gin.Context) {
	data, err := service.GetAllArrangeGroup()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// UpdateStudentsState 更改学生状态
func UpdateStudentsState(c *gin.Context) {
	var par models.ParamStudents
	if err := c.ShouldBind(&par); err != nil {
		zap.L().Error(" c.ShouldBind(&par) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := service.UpdateStudentsState(&par); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// AddArrangeGroup 添加安排组
func AddArrangeGroup(c *gin.Context) {
	var par models.ParamArrangeGroup
	if err := c.ShouldBind(&par); err != nil {
		zap.L().Error("c.ShouldBind(&par) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	if par.Type != "visit" && par.Type != "interview" {
		c.JSON(1001, gin.H{
			"message": "类型参数错误",
		})
		return
	}
	data, err := service.AddArrangeGroup(&par)
	if err != nil {
		zap.L().Error("service.SetArrangeGroup(&par) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// UpdateVisitTime 设置宣讲时间
func UpdateVisitTime(c *gin.Context) {
	var par models.ParamArrange
	if err := c.ShouldBind(&par); err != nil {
		zap.L().Error("c.ShouldBindJSON(par) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err := service.UpdateVisitTime(&par)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// UpdateInterviewTime 设置面试时间
func UpdateInterviewTime(c *gin.Context) {
	var par models.ParamArrange
	if err := c.ShouldBind(&par); err != nil {
		zap.L().Error("c.ShouldBindJSON(par) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err := service.UpdateInterviewTime(&par)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
