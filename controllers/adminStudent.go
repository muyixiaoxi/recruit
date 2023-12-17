package controllers

import (
	"github.com/gin-gonic/gin"
	"recruit/service"
)

// GetStudentInSeven  获取七天内学生
func GetStudentInSeven(c *gin.Context) {
	data, err := service.GetStudentInSeven()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GroupBySpecialty 根据专业分组
func GroupBySpecialty(c *gin.Context) {
	data, err := service.GroupBySpecialty()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
