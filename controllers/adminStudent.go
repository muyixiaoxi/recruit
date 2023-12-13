package controllers

import (
	"github.com/gin-gonic/gin"
	"recruit/service"
)

// GroupBySpecialty 根据专业分组
func GroupBySpecialty(c *gin.Context) {
	data, err := service.GroupBySpecialty()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// UpdateVisitState 同步学生宣讲状况
func UpdateVisitState(c *gin.Context) {
	err := service.UpdateVisitState()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// UpdateInterviewState 同步学生面试状况
func UpdateInterviewState(c *gin.Context) {
	err := service.UpdateInterviewState()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
