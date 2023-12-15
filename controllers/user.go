package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"recruit/models"
	"recruit/pkg/jwt"
	"recruit/service"
)

// GetDetailInfo 获取学生详细信息
func GetDetailInfo(c *gin.Context) {
	var ps models.ParamStudent
	if err := c.ShouldBind(&ps); err != nil {
		zap.L().Error("c.ShouldBind(&ps) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	var s = models.Student{Model: gorm.Model{ID: ps.Id}}
	err := service.GetDetailInfo(&s)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		zap.L().Error("service.GetDetailInfo(openid) failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, s)
}

// ScreenStudents 筛选学生
func ScreenStudents(c *gin.Context) {
	var par models.ParamScreen
	if err := c.ShouldBind(&par); err != nil {
		zap.L().Error("c.ShouldBind(s) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	fmt.Println(par)
	data, err := service.ScreenStudents(&par)
	if err != nil && err != gorm.ErrRecordNotFound {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var user = models.User{}
	if err := c.ShouldBind(&user); err != nil {
		zap.L().Error("err := c.ShouldBind(&user) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	service.UserLogin(&user)
	// 登录失败
	if user.ID == 0 {
		ResponseError(c, CodeInvalidPassword)
		return
	} else {
		service.OnlineUser[user.ID] = make(chan models.InterviewRecord, 5)
		//登陆成功，生成JWT
		token, _ := jwt.GenToken(user.ID)
		user.Token = token
		ResponseSuccess(c, gin.H{
			"token":    user.Token,
			"username": user.Name,
			"role":     user.Role,
		})
	}
}

// GetAllSignUp 获取所有报名学生信息
func GetAllSignUp(c *gin.Context) {
	data, err := service.GetAllSignUp()
	if err != nil {
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}
