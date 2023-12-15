package controllers

// 面试、宣讲、参观安排
import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"recruit/models"
	"recruit/service"
	"strconv"
)

// GetArrangeDetail 获取安排组详细信息
func GetArrangeDetail(c *gin.Context) {
	par := c.Query("id")
	id, _ := strconv.Atoi(par)
	data, err := service.GetArrangeDetail(uint(id))
	if err != nil {
		zap.L().Error("service.GetArrangeDetail(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// DeleteArrange 删除安排组
func DeleteArrange(c *gin.Context) {
	var par models.ParamIds
	if err := c.ShouldBind(&par); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := service.DeleteArrange(par.Id); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseError(c, CodeSuccess)
}

// GetArrangeMenus 获取面试安排菜单
func GetArrangeMenus(c *gin.Context) {
	par := c.Query("type")
	t, _ := strconv.Atoi(par)
	data, err := service.GetArrangeMenus(t)
	if err != nil {
		zap.L().Error("service.GetArrangeMenus() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	return
}

// CancelTime 删除某些学生已有安排组
func CancelTime(c *gin.Context) {
	par := models.ParamCancelArrangeTime{}
	if err := c.ShouldBind(&par); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := service.CancelTime(par); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

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
